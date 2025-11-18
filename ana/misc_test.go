package ana

import (
	"testing"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
