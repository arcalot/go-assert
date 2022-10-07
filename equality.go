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
func Equals[T any](t *testing.T, got T, expected T) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Mismatch, expected: %v, got: %v", expected, got)
	}
}
