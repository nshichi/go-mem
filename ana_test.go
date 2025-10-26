package ana

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func Test_int_and_pointer(t *testing.T) {
	var i int = 123
	var p *int = &i

	fmt.Printf("&i -> %p\n", &i) // 変数 i のアドレス
	fmt.Printf("p  -> %p\n", p)  // 変数　p の値
	fmt.Printf("&p -> %p\n", &p) // 変数 p のアドレス
}

func Test_nil(t *testing.T) {
	var p *int = nil
	fmt.Printf("nil -> %p\n", p)
	fmt.Printf("nil -> %p\n", (*int)(nil))
	fmt.Printf("nil -> %v\n", nil)
}

func Test_1(t *testing.T) {
	s1 := "abc"
	s2 := "abc"
	s3 := string([]rune{'a', 'b', 'c'})
	t.Logf("address of s1 -> %p", unsafe.StringData(s1))
	t.Logf("address of s2 -> %p", unsafe.StringData(s2))
	t.Logf("address of s3 -> %p", unsafe.StringData(s3))
	t.Logf("s1 == s3 -> %v", s1 == s3)
	t.Logf("sizeof(s1) - >%d", unsafe.Sizeof(s1))

	h := "hello world!"
	w := h[6:11]
	t.Logf("address of h \"%s\"; %p", h, unsafe.StringData(h))
	t.Logf("address of w \"%s\"; %p", w, unsafe.StringData(w))

	h = "hello"
	t.Logf("\"%s\" %p", h, unsafe.StringData(h))
	h += " world"
	t.Logf("\"%s\" %p", h, unsafe.StringData(h))
}

func Benchmark1_1(b *testing.B) {
	for b.Loop() {
		s1 := "abc"
		s2 := "abc"
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

func Test_スライス継ぎ足し(t *testing.T) {
	var a []int

	t.Logf("a == nil -> %v", a == nil)
	t.Logf("cap(a), len(a) -> %d, %d", cap(a), len(a))
	// t.Logf("a.Data -> %p", unsafe.SliceData(a))

	a = make([]int, 0)
	t.Logf("a == nil -> %v", a == nil)
	t.Logf("cap(a), len(a) -> %d, %d", cap(a), len(a))
	// t.Logf("a.Data -> %p", unsafe.SliceData(a))

	a = append(a, 0)
	t.Logf("cap(a), len(a), Data -> %d, %d, %p", cap(a), len(a), &a[0])

	a = append(a, 1)
	t.Logf("cap(a), len(a), Data -> %d, %d, %p", cap(a), len(a), &a[0])

	a = append(a, 2)
	t.Logf("cap(a), len(a), Data -> %d, %d, %p", cap(a), len(a), &a[0])

	a = append(a, 3)
	t.Logf("cap(a), len(a), Data -> %d, %d, %p", cap(a), len(a), &a[0])

	a = append(a, 4)
	t.Logf("cap(a), len(a), Data -> %d, %d, %p", cap(a), len(a), &a[0])

	// リセット
	a = a[:0]
	t.Logf("cap(a), len(a) -> %d, %d", cap(a), len(a))
	// t.Logf("a.Data -> %p", &a[0])
}

func BenchmarkXxx(b *testing.B) {
	for b.Loop() {
		var s string
		for range 1_000_000 {
			s += "A"
		}
	}
} /*
goos: windows
goarch: amd64
pkg: go-mem
cpu: AMD Ryzen 7 7800X3D 8-Core Processor
=== RUN   BenchmarkXxx
BenchmarkXxx
BenchmarkXxx-16

	1        29137388100 ns/op       503993288400 B/op        1006323 allocs/op

PASS
ok      go-mem  29.388s
*/

func BenchmarkXxx2(b *testing.B) {
	for b.Loop() {
		tt := make([]string, 0)
		for range 1_000_000 {
			tt = append(tt, "A")
		}
		s := strings.Join(tt, "")
		_ = s
	}
} /*
goos: windows
goarch: amd64
pkg: go-mem
cpu: AMD Ryzen 7 7800X3D 8-Core Processor
=== RUN   BenchmarkXxx2
BenchmarkXxx2
BenchmarkXxx2-16
      45          25018133 ns/op        89025396 B/op         40 allocs/op
PASS
ok      go-mem  1.363s
*/

func Test999(t *testing.T) {
	var a int
	var b int = 456
	var p = new(int)
	t.Logf("&a -> %p", &a)
	t.Logf("&b -> %p", &b)
	t.Logf("p -> %p", p)

	var q = func1(t)
	t.Logf("q -> %p", q)
}

func func1(t *testing.T) *int {
	var a int
	var b int = 456
	var p = new(int)

	t.Logf("&a -> %p", &a)
	t.Logf("&b -> %p", &b)
	t.Logf("p -> %p", p)

	return &b
}

func Test998(t *testing.T) {
	var a int
	var p = new(int)
	t.Logf("&a -> %p", &a)
	t.Logf("p -> %p", p)
}

func TestXxx0(t *testing.T) {
	a := make([]int, 10)
	t.Logf("%d, %d", len(a), cap(a))

	a = a[:0]
	t.Logf("%d, %d", len(a), cap(a))
}
