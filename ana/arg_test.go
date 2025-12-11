package ana

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_takeInt(t *testing.T) {
	var i int = 123
	fmt.Printf("           &i -> %p: %v\n", &i, i)

	takeInt(i)
	fmt.Printf("           &i -> %p: %v\n", &i, i)

	takePInt(&i)
	fmt.Printf("           &i -> %p: %v\n", &i, i)
}

func takeInt(i int) {
	i = 999
	fmt.Printf("takeInt(): &i -> %p: %v\n", &i, i)
}

func takePInt(i *int) {
	*i = 99
	fmt.Printf("takePInt(): i -> %p: %v\n", i, *i)
}

func Test_takeStr(t *testing.T) {
	var s string = "ABC"
	fmt.Printf("           &s -> %p: %p \"%s\"\n", &s, unsafe.StringData(s), s)

	takeStr(s)
	fmt.Printf("           &s -> %p: %p \"%s\"\n", &s, unsafe.StringData(s), s)

	takePStr(&s)
	fmt.Printf("           &s -> %p: %p \"%s\"\n", &s, unsafe.StringData(s), s)
}

func takeStr(s string) {
	s = "zzz"
	fmt.Printf("takeStr(): &s -> %p: %p \"%s\"\n", &s, unsafe.StringData(s), s)
}

func takePStr(s *string) {
	*s = "zz"
	fmt.Printf("takePStr(): s -> %p: %p \"%s\"\n", s, unsafe.StringData(*s), *s)
}

func Test_takeArray(t *testing.T) {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Printf("              &a -> %p: %v\n", &a, a)

	takeArray(a)
	fmt.Printf("              &a -> %p: %v\n", &a, a)

	takePArray(&a)
	fmt.Printf("              &a -> %p: %v\n", &a, a)
}

func takeArray(a [3]int) {
	a[1] = 999
	fmt.Printf("takeArray():  &a -> %p: %v\n", &a, a)
}

func takePArray(a *[3]int) {
	(*a)[1] = 999
	a[2] = 1000
	fmt.Printf("takePArray(): &a -> %p: %v\n", a, *a)
}

func Test_takeSlice(t *testing.T) {
	var s []int = []int{1, 2, 3}
	fmt.Printf("              &s -> %p: %v\n", &s, s)

	takeSlice(s)
	fmt.Printf("              &s -> %p: %v\n", &s, s)

	takePSlice(&s)
	fmt.Printf("              &s -> %p: %v\n", &s, s)
}

func takeSlice(s []int) {
	s[1] = 99

	// fmt.Printf("takeSlice     &s -> %p: %v\n", &s, s)
}

func takePSlice(s *[]int) {
	(*s)[1] = 999
	// s[2] = 1000

	// fmt.Printf("takeSlice     &s -> %p: %v\n", s, *s)
}

func Test_takeMap(t *testing.T) {
	var m = make(map[string]int)
	fmt.Printf("              &m -> %p: %v\n", &m, m)

	takeMap(m)
	fmt.Printf("              &m -> %p: %v\n", &m, m)

	takePMap(&m)
	fmt.Printf("              &m -> %p: %v\n", &m, m)
}

func takeMap(m map[string]int) {
	m["a"] = 999
}

func takePMap(m *map[string]int) {
	(*m)["b"] = 1000
}
