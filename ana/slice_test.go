package ana

import (
	"encoding/json"
	"testing"
	"unsafe"
)

// スライス コピー
func Test_copy_slice(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := s
	s[1] = 999
	t.Logf("s  -> %v", s)
	t.Logf("s1 -> %v", s1)
	t.Logf("s.Data  -> %p", unsafe.SliceData(s))
	t.Logf("s1.Data -> %p", unsafe.SliceData(s1))

	s2 := make([]int, 5)
	n := copy(s2, s)
	t.Logf("copy(s2, s) -> %v", n)
	s2[1] = 888
	t.Logf("s  -> %v", s)
	t.Logf("s2 -> %v", s2)

	t.Logf("s2.Data -> %p", unsafe.SliceData(s2))
}

// 配列 代入
func Test_assign_array(t *testing.T) {
	a1 := [3]int{1, 1, 1}
	a2 := a1
	a1[1] = 999
	t.Logf("a1 -> %v", a1)
	t.Logf("a2 -> %v", a2)
}

// スライス 関数渡し
func Test_arg_slice(t *testing.T) {
	s := []int{1, 1, 1}
	t.Logf("s -> %v", s)

	func(a2 []int) {
		a2[1] = 999
	}(s)
	t.Logf("s' -> %v", s)
}

// 配列 関数コピー渡し
func Test_arg_array(t *testing.T) {
	a := [3]int{1, 1, 1}
	t.Logf("a -> %v", a)

	func(a2 [3]int) {
		a2[1] = 999
	}(a)
	t.Logf("a' -> %v", a)
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

func Test_slicing(t *testing.T) {
	s := []int{1, 2, 3}
	t.Logf("s, cap(s) -> %v, %d", s, cap(s))

	s1 := s[1:2]
	t.Logf("s[1:2], cap(s[1:2]) -> %v, %d", s1, cap(s1))

	a := [5]int{1, 2, 3, 4, 5}
	t.Logf("a, cap(a) -> %v, %d", a, cap(a))

	s3 := a[1:2]
	t.Logf("a[1:2], cap(a[1:2]) -> %v, %d", s3, cap(s3))

	// s4 := s3[0:4]
	// t.Logf("s4, cap(s4) -> %v, %d", s4, cap(s4))

	s5 := a[1:2][0:4]
	t.Logf("a[1:2][0:4], cap(a[1:2][0:4]) -> %v, %d", s5, cap(s5))
}

func Test_json_marshal(t *testing.T) {
	var s []int // 未初期化 = nil スライス

	if s == nil {
		t.Logf("s is nil slice")
	} else {
		t.Logf("s is not nil slice, len(s) -> %d", len(s))
	}
	j0, _ := json.Marshal(s)
	t.Logf(`s0 -> "%s"`, string(j0))

	s = []int{}
	if s == nil {
		t.Logf("s is nil slice")
	} else {
		t.Logf("s is not nil slice, len(s) -> %d", len(s))
	}
	j1, _ := json.Marshal(s)
	t.Logf(`s -> "%s"`, string(j1))

	s = nil
	if s == nil {
		t.Logf("s is nil slice")
	} else {
		t.Logf("s is not nil slice, len(s) -> %d", len(s))
	}
	j2, _ := json.Marshal(s)
	t.Logf(`s -> "%s"`, string(j2))
}

func Test_nil_slice(t *testing.T) {
	eq := map[bool]string{
		false: "!=",
		true:  "==",
	}

	// 未初期化 は nil スライス
	var s []int
	j, _ := json.Marshal(s)
	t.Logf(`no init     %s nil, len(s) -> %d, json -> "%s"`, eq[s == nil], len(s), string(j))

	// 空スライス は nilスライスではない
	s = []int{}
	j, _ = json.Marshal(s)
	t.Logf(`empty slice %s nil, len(s) -> %d, json -> "%s"`, eq[s == nil], len(s), string(j))

	// スライスにnilを代入
	s = nil
	j, _ = json.Marshal(s)
	t.Logf(`assign nil  %s nil, len(s) -> %d, json -> "%s"`, eq[s == nil], len(s), string(j))
}

func Test_slice_shrink(t *testing.T) {
	var s = []int{1, 2, 3, 4, 5}
	t.Logf("s -> %v", s)

	s = s[:0]
	t.Logf("s -> %v", s)

	s = s[:4]
	t.Logf("s -> %v", s)

	// s = s[:6]
	// t.Logf("s -> %v", s)
	// panic: runtime error: slice bounds out of range [:6] with capacity 5 [recovered, repanicked]
}
