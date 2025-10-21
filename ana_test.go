package ana

import (
	"strings"
	"testing"
	"unsafe"
)

func Benchmark1(b *testing.B) {
	for b.Loop() {
		s1 := strings.Repeat("a", 1_000_000)
		s2 := strings.Repeat("b", 1_000_000)
		u := s1 + s2
		_ = u
	}
	// f, err := os.Open("nul")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(f, "%s", u)
}

func Benchmark2(b *testing.B) {
	for b.Loop() {
		const s1 = "strings.Repeat(\"a\", 1_000_000)"
		// const s2 = "strings.Repeat(\"b\", 1_000_000)"
		var u = ""
		for range 100_000 {
			u = u + s1
		}
		_ = u
	}
	// f, err := os.Open("nul")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(f, "%s", u)
}

type R struct {
	A int
	B int32
	C int64
	D float32
	E float64
	F int
	G []int

	P *int
	Q int
	R struct {
		R1 uint
		R2 uint
	}
	S string
}

func Benchmark3(b *testing.B) {
	b.Logf("%d bytes", unsafe.Sizeof(R{}))
	for b.Loop() {
		aa := make([]R, 0)
		for range 1_000 {
			aa = append(aa, R{})
		}
		_ = aa
	}
}
