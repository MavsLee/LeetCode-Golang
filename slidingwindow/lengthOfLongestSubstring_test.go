package slidingwindow

import (
	"testing"
)

func TestLOS(t *testing.T) {
	// to char ints
	testData := []struct {
		s        string
		expected int
	}{
		{
			s:        "",
			expected: 0,
		},
		{
			s:        "a",
			expected: 1,
		},
		{
			s:        "ab",
			expected: 2,
		},
		{
			s:        "abbc",
			expected: 2,
		},
		{
			s:        "abbcdea",
			expected: 5,
		},
	}
	for _, test := range testData {
		if res := LengthOfLongestSubstring(test.s); res != test.expected {
			t.Fatalf("case %s failed, expecte: %v, actual: %v", test.s, test.expected, res)
		}

	}
}
