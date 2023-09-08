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
		//nolint:gocritic,nolintlint
		t.Fatalf("Bad instance. Tested value %v is instance of %T, not %T", got, got, *new(T))
	}
}

type Ordered interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float64 | float32 | ~string
}

// GreaterThan asserts that the got value is greater than the expected minimum.
func GreaterThan[T Ordered](t *testing.T, got T, expectedMin T) {
	t.Helper()
	if got <= expectedMin {
		t.Fatalf("Assertion failed, got value '%v' is not greater than '%v'", got, expectedMin)
	}
}

// LessThan asserts that the got value is greater than the expected minimum.
func LessThan[T Ordered](t *testing.T, got T, expectedMax T) {
	t.Helper()
	if got >= expectedMax {
		t.Fatalf("Assertion failed, got value '%v' is not less than '%v'", got, expectedMax)
	}
}

// GreaterThanEq asserts that the got value is greater than or equal to the expected minimum.
func GreaterThanEq[T Ordered](t *testing.T, got T, expectedMin T) {
	t.Helper()
	if got < expectedMin {
		t.Fatalf("Assertion failed, got value '%v' is not greater than or equal to '%v'", got, expectedMin)
	}
}

// LessThanEq asserts that the got value is greater than or equal to the expected minimum.
func LessThanEq[T Ordered](t *testing.T, got T, expectedMax T) {
	t.Helper()
	if got > expectedMax {
		t.Fatalf("Assertion failed, got value '%v' is not less than or equal to '%v'", got, expectedMax)
	}
}
