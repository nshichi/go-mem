package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	os.Chdir("../ana/")
	cmd := exec.Command("go", "test", "-bench", ".", "-benchmem", "-run=^$")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Printf("running...; %s\n", cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("%s\n", stderr.String())
		return
	}

	filename := "bench-out/" + makeFilename()
	// if _, err := os.Stat(filename); err == nil {
	// 	panic()
	// }
	if err := os.WriteFile(filename, stdout.Bytes(), 0664); err != nil {
		panic(err)
	}

	// results, sysInfo := bench.Parse(stdout.String())

	// fmt.Println("システム情報")
	// fmt.Printf("OS/Arch: %s/%s\n", sysInfo.GOOS, sysInfo.GOARCH)
	// fmt.Printf("CPU: %s\n", sysInfo.CPU)
	// // fmt.Printf("パッケージ: %s\n", sysInfo.Pkg)

	// for _, res := range results {
	// 	fmt.Fprintf(o, "%s-%d: %d ns/op, %d B/op, %d allocs/op\n",
	// 		res.Name,
	// 		res.CPUNum,
	// 		res.NsPerOp,
	// 		res.BytesPerOp,
	// 		res.AllocsPerOp,
	// 		// res.Iters,
	// 		// time.Now().Format("2006-01-02"),
	// 	)
	// }
}

func makeFilename() string {
	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}

	// goos := runtime.GOOS
	// goarch := runtime.GOARCH
	d := time.Now().Format("2006-01-02")

	return fmt.Sprintf("%s %s.txt", d, host)
}
