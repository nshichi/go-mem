package ana

import (
	"testing"
	"time"
	"unsafe"
)

type R struct {
	Int     int
	Pointer *int
	String  string
	Array   [10]int
	Slice   []string
	Time    time.Time // struct
	// Float64 float64
	// Bool    bool
	// Rune    rune
	// Bytes   [10]byte
	// Any     any
	// Map     map[string]string
}

func Test_struct_size(t *testing.T) {
	var r R
	r.String = "abcdefghijklmnopqrstuvwxyz"
	r.Slice = []string{"1", "2", "3", "4", "5"}

	t.Logf("r         -> %p, %d", &r, unsafe.Sizeof(r))
	t.Logf("r.Int     -> %p, %d", &r.Int, unsafe.Sizeof(r.Int))
	t.Logf("r.Pointer -> %p, %d", &r.Pointer, unsafe.Sizeof(r.Pointer))
	t.Logf("r.String  -> %p, %d", &r.String, unsafe.Sizeof(r.String))
	t.Logf("r.Array   -> %p, %d", &r.Array, unsafe.Sizeof(r.Array))
	t.Logf("r.Slice   -> %p, %d", &r.Slice, unsafe.Sizeof(r.Slice))
	t.Logf("r.Time    -> %p, %d", &r.Time, unsafe.Sizeof(r.Time))
}

func Benchmark_struct_append(b *testing.B) {
	for b.Loop() {
		aa := make([]R, 0, 1_000)
		for range 1_000 {
			aa = append(aa, R{})
		}
		_ = aa
	}
}

func Benchmark_pointer_append(b *testing.B) {
	for b.Loop() {
		aa := make([]*R, 0, 1_000)
		for range 1_000 {
			aa = append(aa, &R{})
		}
		_ = aa
	}
}

func Test_struct_empty(t *testing.T) {
	var e0 struct{}
	var e1 struct{}
	var ea [3]struct{}
	t.Logf("&e0, size of e0 -> %p, %d", unsafe.Pointer(&e0), unsafe.Sizeof(e0))
	t.Logf("&e1, size of e1 -> %p, %d", unsafe.Pointer(&e1), unsafe.Sizeof(e1))
	t.Logf("&ea, size of ea -> %p, %d", unsafe.Pointer(&ea), unsafe.Sizeof(ea))

	var es []struct{}
	es = append(es, struct{}{})
	t.Logf("&es, size of es -> %p, %d", unsafe.Pointer(&es), unsafe.Sizeof(es))
	t.Logf("&es[0], size of es[0] -> %p, %d", unsafe.Pointer(&es[0]), unsafe.Sizeof(es[0]))
}
