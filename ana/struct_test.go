package ana

import (
	"testing"
	"time"
	"unsafe"
)

type R struct {
	I     int
	P     *int
	S     string
	Array [10]int
	Slice []string
	Time  time.Time // struct
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

func Test_show_nil(t *testing.T) {
	r := SampleStruct{}
	r.S = "abcdefghijklmnopqrstuvwxyz"
	r.Slice = []int{1, 2, 3, 4, 5}
	t.Logf("sizeof(r) = %d", unsafe.Sizeof(r))
	t.Logf("sizeof(string) = %d", unsafe.Sizeof(r.S))
	t.Logf("sizeof(int) = %d", unsafe.Sizeof(r.I))
	t.Logf("sizeof(pointer) = %d", unsafe.Sizeof(r.P))
	t.Logf("sizeof([10]int) = %d", unsafe.Sizeof(r.Array))
	t.Logf("sizeof(slice) = %d", unsafe.Sizeof(r.Slice))
	t.Logf("sizeof(time.Time) = %d", unsafe.Sizeof(r.Time))

	var p *int = nil
	t.Logf("%p", p)
}
