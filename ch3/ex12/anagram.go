package anagram

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	b := strings.Split(s, "")
	sort.Strings(b)
	return strings.Join(b, "")
}

func AreAnagram(s1 string, s2 string) bool {
	if n := len(s1); n <= 1 {
		return false
	}

	return sortString(s1) == sortString(s2)
}
