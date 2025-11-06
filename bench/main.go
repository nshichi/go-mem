package main

import (
	"bytes"
	"fmt"
	"go-mem/bench/bench"
	"os"
	"os/exec"
)

func main() {
	os.Chdir("../ana/")
	cmd := exec.Command("go", "test", "-bench", ".", "-benchmem", "-run=^$")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Printf("running...; %s\n", "go test -bench . -benchmem -run=^$")
	if err := cmd.Run(); err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("%s\n", stderr.String())
		return
	}

	// os.WriteFile("../benchi/bench.out", stdout.Bytes(), 0664)

	results, sysInfo := bench.Parse(stdout.String())

	fmt.Println("システム情報")
	fmt.Printf("OS/Arch: %s/%s\n", sysInfo.GOOS, sysInfo.GOARCH)
	fmt.Printf("CPU: %s\n", sysInfo.CPU)
	// fmt.Printf("パッケージ: %s\n", sysInfo.Pkg)

	for _, res := range results {
		fmt.Printf("  %s-%d (Iters: %d): %d ns/op, %d B/op, %d allocs/op\n",
			res.Name,
			res.CPUNum,
			res.Iters,
			res.NsPerOp,
			res.BytesPerOp,
			res.AllocsPerOp,
		)
	}
}
