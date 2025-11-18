package ana

import (
	"math/rand/v2"
	"testing"
	"unsafe"
)

func Test_map(t *testing.T) {
	myMap := make(map[string]int)
	myMap["a"] = 1

	mapPtr := unsafe.Pointer(&myMap)
	hmapPtr := (*unsafe.Pointer)(mapPtr) // runtime.hmap
	t.Logf("mapPtr -> %p, *hmapPtr -> %p", mapPtr, *hmapPtr)
}

// cf. 【Go】Mapの内部構造とO(1)のメカニズム https://zenn.dev/smartshopping/articles/5df9c3717e25bd

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
