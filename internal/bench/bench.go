package bench

import (
	"regexp"
	"strconv"
	"strings"
)

type Result struct {
	Name        string
	CPUNum      int
	Iters       int
	NsPerOp     int
	BytesPerOp  int
	AllocsPerOp int
}

type SystemInfo struct {
	GOOS   string
	GOARCH string
	Pkg    string
	CPU    string
}

const reBench = `^(Benchmark[a-zA-Z0-9_/]+)-(\d+)\s+(\d+)\s+([\d.]+)\s+ns/op\s+([\d.]+)\s+B/op\s+([\d.]+)\s+allocs/op`

func Parse(output string) ([]Result, SystemInfo) {
	results := []Result{}
	sysInfo := SystemInfo{}

	re := regexp.MustCompile(reBench)

	lines := strings.SplitSeq(output, "\n")

	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if value, ok := strings.CutPrefix(line, "goos:"); ok {
			sysInfo.GOOS = strings.TrimSpace(value)
		} else if value, ok := strings.CutPrefix(line, "goarch:"); ok {
			sysInfo.GOARCH = strings.TrimSpace(value)
		} else if value, ok := strings.CutPrefix(line, "pkg:"); ok {
			sysInfo.Pkg = strings.TrimSpace(value)
		} else if value, ok := strings.CutPrefix(line, "cpu:"); ok {
			sysInfo.CPU = strings.TrimSpace(value)
		} else {
			matches := re.FindStringSubmatch(line)
			if len(matches) == 7 {
				var r Result
				r.Name = matches[1]
				r.CPUNum, _ = strconv.Atoi(matches[2])
				r.Iters, _ = strconv.Atoi(matches[3])
				r.NsPerOp, _ = strconv.Atoi(matches[4])
				r.BytesPerOp, _ = strconv.Atoi(matches[5])
				r.AllocsPerOp, _ = strconv.Atoi(matches[6])
				results = append(results, r)
			}
		}
	}

	return results, sysInfo
}

func FormatWithCommas(n int) string {
	out := make([]byte, 0, 32) // worst case; MinInt64 -> "-9,223,372,036,854,775,808"

	s := strconv.Itoa(n)
	if s[0] == '-' {
		s = s[1:]
		out = append(out, '-')
	}

	l := len(s)
	i := l % 3
	if i > 0 {
		out = append(out, []byte(s[0:i])...)
	}

	for ; i < l; i += 3 {
		if i > 0 {
			out = append(out, ',')
		}

		out = append(out, []byte(s[i:i+3])...)
	}

	return string(out)
}
