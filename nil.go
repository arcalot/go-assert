package assert

import (
	"reflect"
	"testing"
)

// Nil tests if the provided value is nil. If it is not nil, the assertion fails.
func Nil(t *testing.T, value any) {
	t.Helper()
	if !isNil(value) {
		t.Fatalf("Unexpected non-nil value: %v", value)
	}
}

// NotNil tests that the provided value is not nil. If the value is nil, the assertion fails.
func NotNil(t *testing.T, value any) {
	t.Helper()
	if isNil(value) {
		t.Fatalf("Unexpected nil value")
	}
}

// NotNilR asserts that the input is not nil, and returns it.
func NotNilR[T any](t *testing.T, value T) T {
	t.Helper()
	if isNil(value) {
		t.Fatalf("Unexpected nil value")
	}
	return value
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	nillableKinds := []reflect.Kind{
		reflect.Pointer, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer,
	}
	// Must be a valid value, and must be one that can be nil. If it cannot be nil, then it's not nil.
	if !v.IsValid() || !containsShallow(nillableKinds, v.Kind()) {
		return false
	}
	return v.IsNil()
}
