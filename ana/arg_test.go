package ana

import "testing"

func Test_takeInt(t *testing.T) {
	var i int = 1
	takeInt(i)
	takePInt(&i)
}

func takeInt(i int) {
	i = 2
}

func takePInt(i *int) {
	*i = 2
}

func Test_takeStr(t *testing.T) {
	var s string = "1"
	takeStr(s)
	takePStr(&s)
}

func takeStr(s string) {
	s = "2"
}

func takePStr(s *string) {
	*s = "2"
}

func Test_takeArray(t *testing.T) {
	var a [3]int = [3]int{1, 2, 3}
	takeArray(a)
	takePArray(&a)
}

func takeArray(a [3]int) {
	a[1] = 999
}

func takePArray(a *[3]int) {
	a[1] = 999
}

func Test_takeSlice(t *testing.T) {
	var s []int = []int{1, 2, 3}
	takeSlice(s)
	takePSlice(&s)
}

func takeSlice(s []int) {
	s[1] = 999
}

func takePSlice(s *[]int) {
	(*s)[1] = 999
}
