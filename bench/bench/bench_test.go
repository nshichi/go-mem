package bench

import (
	"math"
	"math/rand"
	"os"
	"testing"
)

func Test_Parse(t *testing.T) {
	bb, err := os.ReadFile("./testdata/bench.out")
	if err != nil {
		t.Fatal(err)
	}

	output := string(bb)
	rr, info := Parse(output)
	_ = rr
	_ = info
	//..

}

type TestData struct {
	N    int
	Want string
}

var testData = []TestData{
	{0, "0"},
	{9, "9"},
	{12, "12"},
	{123, "123"},
	{1_234, "1,234"},
	{12_345, "12,345"},
	{123_456, "123,456"},
	{1_234_567, "1,234,567"},
	{math.MaxInt, "9,223,372,036,854,775,807"},
	{math.MinInt, "-9,223,372,036,854,775,808"},
	{-123, "-123"},
}

func Test_(t *testing.T) {
	for _, test := range testData {
		got := FormatWithCommas(test.N)
		if got != test.Want {
			t.Errorf("mismatched; %s != %s", got, test.Want)
		}
	}
}

func Benchmark_max(b *testing.B) {
	for b.Loop() {
		got := FormatWithCommas(math.MaxInt)
		_ = got
	}
}

func Benchmark_zero(b *testing.B) {
	for b.Loop() {
		got := FormatWithCommas(0)
		_ = got
	}
}

func Benchmark_(b *testing.B) {
	tests := make([]TestData, 0)
	for b.Loop() {
		tests = append(tests, TestData{
			N: rand.Int(),
		})
	}

	b.ResetTimer()

	for _, test := range tests {
		got := FormatWithCommas(test.N)
		_ = got
	}
}

// func Test_FormatWithCommas3(t *testing.T) {
// 	for _, d := range testData {
// 		got := FormatWithCommasOptimized(d.N)
// 		if got != d.Want {
// 			t.Errorf("mismatched; %v != %v", got, d.Want)
// 		}
// 	}
// }

// // FormatWithCommasOptimized は、整数値を3桁区切りのカンマ付き文字列に変換します。
// // 最終的なバイトスライスの長さを事前に計算し、メモリ割り当てとスライス操作を最小限に抑えます。
// func FormatWithCommasOptimized(n int) string {
// 	// 1. 整数を文字列に変換 (strconv.AppendInt を使用して、直接バイトスライスに書き込む)
// 	buf := make([]byte, 0, 20) // 変換用のバッファ（十分なサイズを確保）
// 	isNegative := n < 0

// 	if isNegative {
// 		n = -n // 負の数を正の数に変換
// 	}

// 	s := strconv.AppendInt(buf, int64(n), 10) // バイトスライスに変換結果を格納

// 	// 2. カンマ挿入に必要な最終的な長さを計算
// 	nDigits := len(s)            // 数字部分の長さ
// 	nCommas := (nDigits - 1) / 3 // 必要なカンマの数

// 	// 最終出力の長さ = 負号 (0 or 1) + 数字の長さ + カンマの数
// 	finalLen := nDigits + nCommas
// 	if isNegative {
// 		finalLen++ // 負号の分
// 	}

// 	// 3. 最終結果用のバッファを一度だけ確保
// 	out := make([]byte, finalLen)

// 	// 4. 負号の処理
// 	outIdx := 0
// 	if isNegative {
// 		out[0] = '-'
// 		outIdx++
// 	}

// 	// 5. 後ろから数字をコピーしながらカンマを挿入
// 	sIdx := nDigits - 1 // 元の数字の文字列のインデックス（末尾から）

// 	for i := 0; i < nDigits; i++ {
// 		// 出力バッファの書き込み位置
// 		outWriteIdx := finalLen - 1 - i

// 		// 3桁ごとにカンマを挿入
// 		if i > 0 && i%3 == 0 {
// 			out[outWriteIdx] = ','
// 			outWriteIdx--
// 		}

// 		// 数字をコピー
// 		out[outWriteIdx] = s[sIdx]
// 		sIdx--
// 	}

// 	return string(out)
// }
