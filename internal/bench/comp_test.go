package bench

const sample_line = `Benchmark_mapS-8                             	      10	 106607396 ns/op	111623905 B/op	    8199 allocs/op`

// var benchLineRegex = regexp.MustCompile(
// 	`^(Benchmark[^\s]+)(?:-\d+)?\s+(\d+)\s+([\d\.]+)(ns|us|ms|s)/op\s+([\d\.]+)(B|kB|MB|GB)/op\s+(\d+)\s+allocs/op$`,
// )

// // var re_Benchmark = regexp.MustCompile(`^Benchmark\s+(\d+)\s+(\d+\.\d+)\sns/op\s+(\d+)\sB/op\s(\d+)\sallocs/op`)
// var re_Benchmark = regexp.MustCompile(`^(Benchmark.*)-(\d+)\s+(\d+)\s+(.*) ns/op\s+(.*) B/op\s+(.*) allocs/op(.*)`)

// type BenchLine struct {
// 	Text string
// 	FuncName string
// 	Cores int
// 	Iter int
// 	ns float32
// 	B int
// 	Allocs int
// }

// func Test_RegExp(t *testing.T) {
// 	match := re_Benchmark.FindStringSubmatch(sample_line)
// 	t.Logf("match = %v", match)
// }
