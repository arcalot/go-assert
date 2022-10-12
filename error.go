package assert

import (
	"testing"
)

// NoError checks if there was no error provided.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// NoErrorR checks if there was no error provided and returns a value.
func NoErrorR[T any](t *testing.T) func(T, error) T {
	t.Helper()
	return func(r T, err error) T {
		t.Helper()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		return r
	}
}

// Error checks if there was an error provided.
func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("No error returned")
	}
}

// ErrorR checks if there was an error provided.
func ErrorR[T any](t *testing.T) func(T, error) {
	t.Helper()
	return func(r T, err error) {
		t.Helper()
		if err == nil {
			t.Fatalf("No error returned")
		}
	}
}
