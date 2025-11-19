package ana

import (
	"math/rand/v2"
	"testing"
	"unsafe"
)

func Test_map(t *testing.T) {
	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1

	mp := unsafe.Pointer(&m)
	hmapPtr := (*unsafe.Pointer)(mp) // runtime.hmap
	t.Logf("mp -> %p, *hmapPtr -> %p", mp, *hmapPtr)
}

// cf. 【Go】Mapの内部構造とO(1)のメカニズム https://zenn.dev/smartshopping/articles/5df9c3717e25bd

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

func Benchmark_mapI(b *testing.B) {
	keys := make([]uint64, 0)
	for range 1_000_000 {
		var m uint64 = 1
		for range 10 {
			m *= uint64(len(letters))
		}
		keys = append(keys, rand.Uint64N(m))
	}

	b.ResetTimer()

	for b.Loop() {
		m := make(map[uint64]struct{})
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

// map is reference, can be modified with no pointer
func Test_map_add(t *testing.T) {
	var m = make(map[string]int)
	m["A"] = 123
	addmap(m, "B", 999)
	t.Logf("m -> %v", m)
}

func addmap(m map[string]int, key string, value int) {
	m[key] = value
}
