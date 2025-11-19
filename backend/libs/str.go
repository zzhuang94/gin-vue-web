package libs

import (
	"encoding/json"
	"strings"
)

func SplitLines(str string) []string {
	ans := make([]string, 0)
	m := make(map[string]bool)
	for _, line := range strings.Split(str, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, ok := m[line]; ok {
			continue
		}
		m[line] = true
		ans = append(ans, line)
	}
	return ans
}

func ParseAndFormatJson(str string) (string, error) {
	var v any
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		return str, err
	}
	bs, err := json.Marshal(v)
	return string(bs), err
}
