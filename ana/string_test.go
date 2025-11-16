package ana

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

// 文字列
func Test_string(t *testing.T) {
	// 代入
	var s = "hello world"
	var s1 = s
	fmt.Printf("s  -> %p: %s\n", unsafe.StringData(s), s)
	fmt.Printf("s1 -> %p: %s\n", unsafe.StringData(s1), s1)

	// スライシング
	var h = s[:5]
	var w = s[6:]
	p2 := unsafe.Pointer(unsafe.StringData(h))
	p3 := unsafe.Pointer(unsafe.StringData(w))
	fmt.Printf("h -> %p: \"%s\"\n", p2, h)
	fmt.Printf("w -> %p: \"%s\"\n", p3, w)
	fmt.Printf("%p - %p -> %v\n", p3, p2, uintptr(p3)-uintptr(p2))

	// s と同じメモリ (おそらく最適化による)
	var s2 = "hello" + " " + "world"
	fmt.Printf("s2 -> %p: \"%s\"\n", unsafe.StringData(s2), s2)

	// s と別メモリでおなじ値
	var s3 = strings.Join([]string{"hello", " ", "world"}, "")
	fmt.Printf("s5 -> %p: \"%s\"\n", unsafe.StringData(s3), s3)

	// 文字列比較
	t.Logf("s == s1 -> %v", s == s1)
	t.Logf("s == s3 -> %v", s == s3)

	// スライシング (拡張) 実行時エラー
	// var s4 = s2[0:11]
	// panic: runtime error: slice bounds out of range [:11] with length 5 [recovered, repanicked]

	// 文字列はイミュータブル コンパイルエラー
	// s0[5] = ' '

	bb := []byte(s)
	t.Logf("-> %s", hex.EncodeToString(bb))

	x := ""
	for i, b := range bb {
		_ = i
		x += fmt.Sprintf("%02X ", b)
	}
	t.Log(x)
}

const (
	APPEND_TIMES = 1_000_000
	SAMPLE_TEXT  = "A"
)

// 文字列継ぎ足し
func Benchmark_string_append1(b *testing.B) {
	for b.Loop() {
		var s = ""
		for range APPEND_TIMES {
			s += SAMPLE_TEXT
			// b.Logf("%p", unsafe.StringData(s))
		}

		b.Logf("%p", unsafe.StringData(s))
		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し
func Test_string_append(t *testing.T) {
	const APPEND_TIMES = 10
	var s string = ""
	t.Logf("%p, %d", unsafe.StringData(s), len(s))
	for range APPEND_TIMES {
		s += SAMPLE_TEXT
		t.Logf("%p, %d", unsafe.StringData(s), len(s))
	}
}

// 文字列継ぎ足し strings.Join()版
func Benchmark_string_append_strings_Join(b *testing.B) {
	for b.Loop() {
		var ss = make([]string, 0)
		for range APPEND_TIMES {
			ss = append(ss, SAMPLE_TEXT)
		}
		s := strings.Join(ss, "")

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し []byte 版
func Benchmark_string_append_bytes(b *testing.B) {
	for b.Loop() {
		bb := make([]byte, 0)

		for range APPEND_TIMES {
			bb = append(bb, []byte(SAMPLE_TEXT)...)
		}

		s := string(bb)

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し []byte 版 maxLength指定
func Benchmark_string_append_bytes_maxLength(b *testing.B) {
	for b.Loop() {
		maxLength := APPEND_TIMES * len(SAMPLE_TEXT)
		bb := make([]byte, 0, maxLength)

		for range APPEND_TIMES {
			bb = append(bb, []byte(SAMPLE_TEXT)...)
		}

		s := string(bb)

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し unsafe.String()
func Benchmark_string_append_unsafe_String(b *testing.B) {
	for b.Loop() {
		maxLength := APPEND_TIMES * len(SAMPLE_TEXT)
		bb := make([]byte, 0, maxLength)

		for range APPEND_TIMES {
			bb = append(bb, []byte(SAMPLE_TEXT)...)
		}

		s := unsafe.String(&bb[0], len(bb))

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し bytes.Buffer 版
func Benchmark_string_append_bytes_Buffer(b *testing.B) {
	for b.Loop() {
		var bb bytes.Buffer

		for range APPEND_TIMES {
			bb.WriteString(SAMPLE_TEXT)
		}

		s := bb.String()

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列の比較 (別のバイト配列)
func Benchmark_string_compare2(b *testing.B) {
	s1 := strings.Repeat("A", 1_000_000)
	s3 := strings.Repeat("A", 1_000_000)

	for b.Loop() {
		if s1 != s3 {
			b.Errorf("")
		}
	}
}

// 文字列の比較 (同じバイト配列を共有する)
func Benchmark_string_compare1(b *testing.B) {
	s1 := strings.Repeat("A", 1_000_000)
	s2 := s1

	for b.Loop() {
		if s1 != s2 {
			b.Errorf("")
		}
	}
}

func Benchmark_string_compare4(b *testing.B) {
	s1 := strings.Repeat("A", 1_000_000)
	ss := make([]string, 0)
	for range b.N {
		ss = append(ss, s1)
	}
	b.ResetTimer()

	for i := range b.N {
		if s1 != ss[i] {
			b.Errorf("")
		}
	}
}

func compare1(s1, s2 string) bool {
	return s1 == s2
}
