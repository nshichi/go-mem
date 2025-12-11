package ana

import (
	"bytes"
	"strings"
	"testing"
	"unsafe"
)

const (
	APPEND_TIMES = 100_000
	SAMPLE_TEXT  = "0123456789"
	// SAMPLE_TEXT  = "ABCDEFGHIJ"
	// APPEND_TIMES = 1_000_000
	// SAMPLE_TEXT  = "A"
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

// 文字列継ぎ足し unsafe.Pointer()
func Benchmark_string_append_unsafe_Pointer(b *testing.B) {
	for b.Loop() {
		maxLength := APPEND_TIMES * len(SAMPLE_TEXT)
		bb := make([]byte, 0, maxLength)

		for range APPEND_TIMES {
			bb = append(bb, []byte(SAMPLE_TEXT)...)
		}

		s := *(*string)(unsafe.Pointer(&bb))

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し bytes.Buffer 版
func Benchmark_string_append_bytes_Buffer(b *testing.B) {
	for b.Loop() {
		var bb bytes.Buffer

		maxLength := APPEND_TIMES * len(SAMPLE_TEXT)
		bb.Grow(maxLength)

		for range APPEND_TIMES {
			bb.WriteString(SAMPLE_TEXT)
		}

		s := bb.String()

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

// 文字列継ぎ足し strings.Builder 版
func Benchmark_string_append_strings_Builder(b *testing.B) {
	for b.Loop() {
		var sb strings.Builder

		maxLength := APPEND_TIMES * len(SAMPLE_TEXT)
		sb.Grow(maxLength)

		for range APPEND_TIMES {
			sb.WriteString(SAMPLE_TEXT)
		}

		s := sb.String()

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}

func Benchmark_string_append_strings_Builder0(b *testing.B) {
	for b.Loop() {
		var sb strings.Builder

		// maxLength := APPEND_TIMES * len(SAMPLE_TEXT)
		// sb.Grow(maxLength)

		for range APPEND_TIMES {
			sb.WriteString(SAMPLE_TEXT)
		}

		s := sb.String()

		if len(s) != APPEND_TIMES*len(SAMPLE_TEXT) {
			b.Errorf("len(s) mismatched")
		}
	}
}
