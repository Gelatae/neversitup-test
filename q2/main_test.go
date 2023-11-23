package main

import (
	"reflect"
	"testing"
)

func TestPremutation(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"", []string{""}},
		{"a", []string{"a"}},
		{"abcd", []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}},
		{"xy", []string{"xy", "yx"}},
		{"123", []string{"123", "132", "213", "231", "312", "321"}},
		{"aaa", []string{"aaa"}},
		{"xyz", []string{"xyz", "xzy", "yxz", "yzx", "zxy", "zyx"}},
		{"ab", []string{"ab", "ba"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
		// Add more test cases here
	}

	for _, test := range tests {
		result := GetPermutation(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
