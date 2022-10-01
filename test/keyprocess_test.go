package test

import (
	"testing"
)

type addTest struct {
	input    string
	expected bool
}

var addTests = []addTest{
	addTest{"497b94cde9b006bc414f19af515a5462266704316e6d44d5cadaf4194cdcf5fa", false},
}

func TestProcessKey(t *testing.T) {
	for _, test := range addTests {
		if CheckKey(test.input) != test.expected {
			t.Errorf("Key %s shows as found but should not have been", test.input)
		}
	}
}
