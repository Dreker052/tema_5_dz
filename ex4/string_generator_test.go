package ex4

import (
	"strings"
	"testing"
	"testing/quick"
)

func TestGenerateString(t *testing.T) {
	property := func(n int) bool {
		s := GenerateString(n)
		return strings.HasPrefix(s, "your value is: ")
	}

	if err := quick.Check(property, nil); err != nil {
		t.Errorf("Property test failed: %v", err)
	}
}
