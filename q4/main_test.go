package main

import (
	"reflect"
	"testing"
)

func TestCountSmileys(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1},
		{[]string{}, 0},
		{[]string{" ;]", ":) ", ";~)"}, 1},
		// Add more test cases here
	}

	for _, test := range tests {
		result := CountSmileys(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
