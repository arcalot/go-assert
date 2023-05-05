package assert

import (
	"reflect"
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

// Equals checks if the values are equal (deeply).
func Equals[T any](t *testing.T, got T, expected T) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Mismatch, expected: %v, got: %v", expected, got)
	}
}

// InstanceOf checks if the value got is equal to the type specified.
func InstanceOf[T any](t *testing.T, got any) {
	_, ok := got.(T)
	if !ok {
		//nolint
		t.Fatalf("Bad instance. Tested value %v is instance of %T, not %T", got, got, *new(T))
	}
}
