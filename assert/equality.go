package assert

import (
	"strings"
	"testing"
)

// Contains checks if the specified substring is found in data.
func Contains[T ~string, K ~string](t *testing.T, data T, substring K) {
	t.Helper()
	if !strings.Contains(
		string(data),
		string(substring),
	) {
		t.Fatalf("Expected substring '%s' not found in '%s'", substring, data)
	}
}

// Contains checks if the specified substring is found in data.
func Equals[T comparable](t *testing.T, a T, b T) {
	t.Helper()
	if a != b {
		t.Fatalf("Espected values '%v' and '%v' to match.", a, b)
	}
}
