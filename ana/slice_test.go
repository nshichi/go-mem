package ana

import "testing"

// スライス コピー
func Test_copy_slice(t *testing.T) {
	a1 := []int{1, 1, 1}
	a2 := a1
	a1[1] = 999
	t.Logf("a1 -> %v", a1)
	t.Logf("a2 -> %v", a2)
}

// スライス 関数渡し
func Test_arg_slice(t *testing.T) {
	a1 := []int{1, 1, 1}
	func(a2 []int) {
		a2[1] = 999

		t.Logf("a1 -> %v", a1)
		t.Logf("a2 -> %v", a2)
	}(a1)
}

// 配列 コピー
func Test_copy_array(t *testing.T) {
	a1 := [3]int{1, 1, 1}
	a2 := a1
	a1[1] = 999
	t.Logf("a1 -> %v", a1)
	t.Logf("a2 -> %v", a2)
}

// 配列 関数コピー渡し
func Test_arg_array(t *testing.T) {
	a1 := [3]int{1, 1, 1}
	func(a2 [3]int) {
		a2[1] = 999

		t.Logf("a1 -> %v", a1)
		t.Logf("a2 -> %v", a2)
	}(a1)
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
