package main

import (
	"bufio"
	"fmt"
	"go-mem/internal/bench"
	"os"
	"path/filepath"
	"strings"
)

// BenchmarkResult は単一のベンチマークの結果を保持します。
// type BenchmarkResult struct {
// 	Runs        int
// 	NsPerOp     float64
// 	BytesPerOp  float64
// 	AllocsPerOp int
// 	FileName    string
// }

// BenchmarkData はベンチマーク名とその結果のマップです。
// キー: ベンチマーク名 (例: "Benchmark_mapS")
// 値: そのベンチマークのファイルごとの結果のスライス
// var BenchmarkData = make(map[string][]BenchmarkResult)

// ベンチマーク結果行を解析するための正規表現
// 例: "Benchmark_mapS-8 10 106607396 ns/op 111623905 B/op 8199 allocs/op"
// var benchLineRegex = regexp.MustCompile(
// 	`^(Benchmark_[^\s]+)(?:-\d+)?\s+(\d+)\s+([\d\.]+)(ns|us|ms|s)/op\s+([\d\.]+)(B|kB|MB|GB)/op\s+(\d+)\s+allocs/op$`,
// )

func main() {
	const path = "../ana/bench-out"
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	bfiles := make(map[string]map[string]*bench.BenchLine)
	fmap := make(map[string]struct{})
	for _, d := range dir {
		path := filepath.Join(path, d.Name())
		bf, err := bench.ParseFile(path)
		if err != nil {
			// fmt.Printf("エラー: ファイル %s の解析中に問題が発生しました: %v\n", filename, err)
			continue
		}
		bfiles[d.Name()] = bf
		for k := range bf {
			fmap[k] = struct{}{}
		}
	}

	// if len(bb) == 0 {
	// 	fmt.Println("解析されたベンチマーク結果が見つかりませんでした。")
	// 	return
	// }

	// printComparison()
	printComparing(bfiles, fmap)
}

func printComparing(bfiles map[string]map[string]*bench.BenchLine, fmap map[string]struct{}) error {
	o, err := os.Create("benchmark.csv")
	if err != nil {
		panic(err)
	}
	defer o.Close()

	wr := bufio.NewWriter(o)

	fmt.Fprintf(wr, "%s,", "machine")
	for fname := range fmap {
		fmt.Fprintf(wr, "%s,", fname)
	}
	fmt.Fprintf(wr, "\n")

	for file, bf := range bfiles {
		ss := strings.Split(file, "_")
		fmt.Fprintf(wr, "%s,", ss[0])
		for fname := range fmap {
			if b, ok := bf[fname]; ok {
				fmt.Fprintf(wr, "%f,", b.NsPerOp)
			} else {
				fmt.Fprintf(wr, ",")
			}
		}
		fmt.Fprintf(wr, "\n")
	}

	return wr.Flush()
}

// printComparison は集計された結果をテーブル形式で出力します。
// func printComparison() {

// 	// ベンチマーク名をソートして出力順を安定させる
// 	var benchNames []string
// 	for name := range BenchmarkData {
// 		benchNames = append(benchNames, name)
// 	}
// 	// 標準ライブラリの sort は使用していませんが、ベンチマーク名順に処理する準備として配列に格納
// 	// sort.Strings(benchNames)

// 	for _, name := range benchNames {
// 		results := BenchmarkData[name]
// 		if len(results) < 2 {
// 			// 比較対象が1つしかない場合はスキップまたは単独で出力する
// 			// ここでは比較が目的なのでスキップします
// 			// fmt.Printf("\n### %s (比較対象不足)\n", name)
// 			continue
// 		}

// 		fmt.Printf("\n### 🚀 %s\n", name)

// 		fmt.Printf("| ファイル名 | %15s | %15s | %15s |\n", "ns/op", "B/op", "allocs/op")
// 		fmt.Printf("| %-10s | %15s | %15s | %15s |\n", strings.Repeat("-", 10), strings.Repeat("-", 15), strings.Repeat("-", 15), strings.Repeat("-", 15))

// 		// 最初の結果を基準として設定
// 		baseResult := results[0]

// 		// 全ての結果を出力
// 		for i, result := range results {
// 			if i == 0 {
// 				// 基準の結果 (ファイル1)
// 				fmt.Printf("| %-10s | %14.2f ns | %13.2f B | %13d allocs |\n",
// 					result.FileName,
// 					result.NsPerOp,
// 					result.BytesPerOp,
// 					result.AllocsPerOp)
// 			} else {
// 				// 比較対象の結果 (ファイル2以降)
// 				nsDiff := (result.NsPerOp/baseResult.NsPerOp)*100.0 - 100.0
// 				bytesDiff := (result.BytesPerOp/baseResult.BytesPerOp)*100.0 - 100.0
// 				allocsDiff := (float64(result.AllocsPerOp)/float64(baseResult.AllocsPerOp))*100.0 - 100.0

// 				fmt.Printf("| %-10s | %14.2f ns (%+5.1f%%) | %13.2f B (%+5.1f%%) | %13d allocs (%+5.1f%%) |\n",
// 					result.FileName,
// 					result.NsPerOp, nsDiff,
// 					result.BytesPerOp, bytesDiff,
// 					result.AllocsPerOp, allocsDiff)
// 			}
// 		}
// 	}
// }
