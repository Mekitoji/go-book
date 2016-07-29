package comma

import "strings"

func Comma(s string) string {
	var prefix string
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		prefix = string(s[0])
		s = s[1:]
	}
	i := strings.Index(s, ".")
	var suffix string
	if i != -1 {
		suffix = s[i:]
		s = s[:i]
	}
	n := len(s)
	if n <= 3 {
		return prefix + s + suffix
	}
	return prefix + Comma(s[:n-3]) + "," + s[n-3:] + suffix
}
