package libs

func UniqSlice(strs []string) []string {
	ans := make([]string, 0)
	m := make(map[string]bool)
	for _, str := range strs {
		if _, ok := m[str]; ok {
			continue
		}
		m[str] = true
		ans = append(ans, str)
	}
	return ans
}
