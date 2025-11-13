package ana

import (
	"math/rand/v2"
	"testing"
	"time"
)

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
		// b.Logf("len(m) = %d", len(m))
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
		// b.Logf("len(m) = %d", len(m))
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
		// b.Logf("len(m) = %d", len(m))
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

func TestXxx0(t *testing.T) {
	a := make([]int, 10)
	t.Logf("%d, %d", len(a), cap(a))

	a = a[:0]
	t.Logf("%d, %d", len(a), cap(a))
}

func Test_ret_ptr(t *testing.T) {
	p := ret_ptr()
	t.Logf("p, *p -> %p, %d", p, *p)
}

func ret_ptr() *int {
	var a = 999
	return &a
}
