package ana

import (
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
