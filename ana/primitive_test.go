package ana

import (
	"fmt"
	"testing"
	"unsafe"
)

// 整数とポインタ
func Test_int_and_pointer(t *testing.T) {
	var i int = 123
	var p *int = &i

	fmt.Printf("&i -> %p\n", &i) // 変数 i のアドレス
	fmt.Printf(" p -> %p\n", p)  // 変数 p の値
	fmt.Printf("&p -> %p\n", &p) // 変数 p のアドレス
}

// ローカル変数へのポインタを返す関数
func Test_return_pointer(t *testing.T) {
	xp := locptr()
	t.Logf("xp, *xp -> %p, %d", xp, *xp)
	if xp == nil || *xp != 456 {
		t.Errorf("*xp %d must be %d", *xp, 456)
	}
}

func locptr() *int {
	var x int
	x = 456
	return &x
}

// nilポインタ
func Test_nil(t *testing.T) {
	var p *int = nil
	fmt.Printf("nil -> %p\n", p)
	fmt.Printf("nil -> %p\n", (*int)(nil))
	fmt.Printf("nil -> %v\n", nil)

	// *p = 999
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]
}

func Test_new(t *testing.T) {
	var a int
	var p = new(int)
	t.Logf("&a -> %p", &a)
	t.Logf("p -> %p", p)
}

func Test_any(t *testing.T) {
	var a any
	a = 123
	t.Logf("sizeof(a) 123- >%d", unsafe.Sizeof(a))

	a = "abc"
	t.Logf("sizeof(a) \"abc\"- >%d", unsafe.Sizeof(a))

	a = R{}
	t.Logf("sizeof(a) struct R- >%d", unsafe.Sizeof(a))
}
