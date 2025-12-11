package bench

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re_Benchmark = regexp.MustCompile(`^Benchmark(.*)-(\d+)\s+(\d+)\s+(.*) ns/op\s+(.*) B/op\s+(.*) allocs/op(.*)`)

type BenchLine struct {
	FuncName    string
	Cores       int
	Runs        int
	NsPerOp     float64
	BytesPerOp  int64
	AllocsPerOp int64
}

func ParseLine(line string) (*BenchLine, error) {
	match := re_Benchmark.FindStringSubmatch(line)
	if len(match) < 7 {
		return nil, fmt.Errorf("parse error")
	}

	cores, _ := strconv.ParseInt(match[2], 10, 32)
	iter, _ := strconv.ParseInt(match[3], 10, 32)
	ns, _ := strconv.ParseFloat(match[4], 64)
	bytes, _ := strconv.ParseInt(match[5], 10, 64)
	allocs, _ := strconv.ParseInt(match[6], 10, 64)
	b := BenchLine{
		FuncName:    match[1],
		Cores:       int(cores),
		Runs:        int(iter),
		NsPerOp:     ns,
		BytesPerOp:  bytes,
		AllocsPerOp: allocs,
	}
	return &b, nil
}

// parseFile は指定されたファイルからベンチマーク結果を抽出します。
func ParseFile(fileName string) (map[string]*BenchLine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bb := make(map[string]*BenchLine)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// BENCH: や PASS, ok などの行はスキップ
		if strings.HasPrefix(line, "goos:") ||
			strings.HasPrefix(line, "goarch:") ||
			strings.HasPrefix(line, "pkg:") ||
			strings.HasPrefix(line, "cpu:") ||
			strings.HasPrefix(line, "--- BENCH:") ||
			strings.HasPrefix(line, "PASS") ||
			strings.HasPrefix(line, "ok") {
			continue
		}

		// ベンチマーク結果の行を解析

		b, err := ParseLine(line)
		if err == nil {
			bb[b.FuncName] = b
		}
		// match := benchLineRegex.FindStringSubmatch(strings.TrimSpace(line))
		// if len(match) == 8 {
		// 	benchName := match[1]
		// 	runs, _ := strconv.Atoi(match[2])

		// 	// ns/op の値と単位を処理
		// 	nsPerOp, _ := strconv.ParseFloat(match[3], 64)
		// 	unitNs := match[4]
		// 	nsPerOp = convertToNs(nsPerOp, unitNs)

		// 	// B/op の値と単位を処理
		// 	bytesPerOp, _ := strconv.ParseFloat(match[5], 64)
		// 	unitBytes := match[6]
		// 	bytesPerOp = convertToBytes(bytesPerOp, unitBytes)

		// 	allocsPerOp, _ := strconv.Atoi(match[7])

		// 	result := BenchmarkResult{
		// 		Runs:        runs,
		// 		NsPerOp:     nsPerOp,
		// 		BytesPerOp:  bytesPerOp,
		// 		AllocsPerOp: allocsPerOp,
		// 		FileName:    fileName,
		// 	}
		// 	BenchmarkData[benchName] = append(BenchmarkData[benchName], result)
		// }
	}

	return bb, scanner.Err()
}
