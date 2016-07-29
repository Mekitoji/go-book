package anagram

import "testing"

func TestAreAnagrams(t *testing.T) {
	var tests = []struct {
		first    string
		second   string
		expected bool
	}{
		{"", "", false},
		{"a", "a", false},
		{"baker", "break", true},
		{"farmer", "framer", true},
	}

	for _, test := range tests {
		if actual := AreAnagram(test.first, test.second); actual != test.expected {
			t.Errorf("AreAnagram(%q, %q) expected %b, got %b", test.first, test.second, test.expected, actual)
		}
	}
}
