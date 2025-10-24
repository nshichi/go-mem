package ana

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"
	"unsafe"
)

const ABC = "abc"

func Test_1(t *testing.T) {
	s1 := "abc"
	s2 := ABC
	s3 := string([]rune{'a', 'b', 'c'})
	t.Logf("%p", unsafe.StringData(s1))
	t.Logf("%p", unsafe.StringData(s2))
	t.Logf("%p", unsafe.StringData(s3))

	t.Logf("s1 == s3 -> %v", s1 == s3)
}

func Benchmark1_1(b *testing.B) {
	for b.Loop() {
		s1 := "abc"
		s2 := ABC
		s3 := string([]rune{'a', 'b', 'c'})
		b.Logf("%p", unsafe.StringData(s1))
		b.Logf("%p", unsafe.StringData(s2))
		b.Logf("%p", unsafe.StringData(s3))

		b.Logf("s1 == s3 -> %v", s1 == s3)
	}
}

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

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 91,592,258
func Benchmark_mapS(b *testing.B) {
	keys := make([]string, 0)
	for range 1_000_000 {
		// l := 1 + rand.IntN(10)
		key := ""
		for range 10 {
			key = key + string(letters[rand.IntN(len(letters))])
		}
		keys = append(keys, key)
	}

	b.ResetTimer()

	for b.Loop() {
		m := make(map[string]struct{})
		for _, key := range keys {
			m[key] = struct{}{}
		}
		b.Logf("len(m) = %d", len(m))
	}
}

// 53,606,432
func Benchmark_mapI(b *testing.B) {
	keys := make([]int, 0)
	for range 1_000_000 {
		m := 1
		for range 10 {
			m *= len(letters)
		}
		keys = append(keys, rand.IntN(m))
	}

	b.ResetTimer()

	for b.Loop() {
		m := make(map[int]struct{})
		for _, key := range keys {
			m[key] = struct{}{}
		}
		b.Logf("len(m) = %d", len(m))
	}
}

type Key struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	F byte
	G byte
	H byte
	I byte
	J byte
}

func Benchmark_mapX(b *testing.B) {
	keys := make([]Key, 0)
	for range 1_000_000 {
		key := Key{}
		key.A = byte(rand.IntN(len(letters)))
		key.B = byte(rand.IntN(len(letters)))
		key.C = byte(rand.IntN(len(letters)))
		key.D = byte(rand.IntN(len(letters)))
		key.E = byte(rand.IntN(len(letters)))
		key.F = byte(rand.IntN(len(letters)))
		key.G = byte(rand.IntN(len(letters)))
		key.H = byte(rand.IntN(len(letters)))
		key.I = byte(rand.IntN(len(letters)))
		key.J = byte(rand.IntN(len(letters)))
		keys = append(keys, key)
	}

	b.ResetTimer()

	for b.Loop() {
		m := make(map[Key]struct{})
		for _, key := range keys {
			m[key] = struct{}{}
		}
		b.Logf("len(m) = %d", len(m))
	}
}

type SampleStruct struct {
	I     int
	S     string
	P     *int
	Array [10]int
	Slice []int
	Time  time.Time
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
