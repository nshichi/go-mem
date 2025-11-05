package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type BenchmarkResult struct {
	Name        string // ベンチマーク名 (例: Benchmark_mapS)
	CPUNum      int    // CPU数 (例: 16)
	Iters       int    // 実行回数 (b.N)
	NsPerOp     int    // 1操作あたりのナノ秒 (ns/op)
	BytesPerOp  int    // 1操作あたりのメモリ確保量 (B/op)
	AllocsPerOp int    // 1操作あたりのアロケーション回数 (allocs/op)
}

type SystemInfo struct {
	GOOS   string
	GOARCH string
	// Pkg    string
	CPU string
}

func main() {
	// 1. go test コマンドの実行
	os.Chdir("../ana/")
	cmd := exec.Command("go", "test", "-bench", ".", "-benchmem", "-run=^$")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Println(">> go test -bench . -benchmem -run=^$ を実行中...")

	err := cmd.Run()
	if err != nil {
		// コマンド実行エラー (exit status 1 など) が発生した場合
		fmt.Printf("コマンド実行エラー: %v\n", err)
		fmt.Printf("標準エラー出力:\n%s\n", stderr.String())
		return
	}

	// 2. 出力結果のパース
	results, sysInfo := parseBench(stdout.String())

	// 3. 結果の表示
	fmt.Println("\n==================================")
	fmt.Println("🛠️ システム情報")
	fmt.Printf("OS/Arch: %s/%s\n", sysInfo.GOOS, sysInfo.GOARCH)
	fmt.Printf("CPU: %s\n", sysInfo.CPU)
	// fmt.Printf("パッケージ: %s\n", sysInfo.Pkg)
	fmt.Println("==================================")

	fmt.Println("ベンチマーク結果")
	for _, res := range results {
		fmt.Printf(
			"  %s-%d (Iters: %d): %d ns/op, %d B/op, %d allocs/op\n",
			res.Name,
			res.CPUNum,
			res.Iters,
			res.NsPerOp,
			res.BytesPerOp,
			res.AllocsPerOp,
		)
	}
}

const reBench = `^(Benchmark[a-zA-Z0-9_/]+)-(\d+)\s+(\d+)\s+([\d.]+)\s+ns/op\s+([\d.]+)\s+B/op\s+([\d.]+)\s+allocs/op`

func parseBench(output string) ([]BenchmarkResult, SystemInfo) {
	results := []BenchmarkResult{}
	sysInfo := SystemInfo{}

	re := regexp.MustCompile(reBench)

	lines := strings.SplitSeq(output, "\n")

	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 実行環境情報 (goos, goarch, cpu) のパース
		if after, ok := strings.CutPrefix(line, "goos:"); ok {
			sysInfo.GOOS = strings.TrimSpace(after)
		} else if after0, ok0 := strings.CutPrefix(line, "goarch:"); ok0 {
			sysInfo.GOARCH = strings.TrimSpace(after0)
		} else if after2, ok2 := strings.CutPrefix(line, "cpu:"); ok2 {
			sysInfo.CPU = strings.TrimSpace(after2)
		} else {
			matches := re.FindStringSubmatch(line)
			if len(matches) == 7 {
				res := BenchmarkResult{
					Name: matches[1],
				}

				res.CPUNum, _ = strconv.Atoi(matches[2])
				res.Iters, _ = strconv.Atoi(matches[3])
				res.NsPerOp, _ = strconv.Atoi(matches[4])
				res.BytesPerOp, _ = strconv.Atoi(matches[5])
				res.AllocsPerOp, _ = strconv.Atoi(matches[6])

				results = append(results, res)
			}
		}
	}

	return results, sysInfo
}
