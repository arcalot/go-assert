package assert

import (
	"reflect"
	"testing"
)

// Nil tests if the provided value is nil. If it is not nil, the assertion fails.
func Nil(t *testing.T, value any) {
	t.Helper()
	v := reflect.ValueOf(value)
	if v.IsValid() && (v.Kind() != reflect.Pointer || !v.IsNil()) {
		t.Fatalf("Unexpected non-nil value: %v", value)
	}
}

// NotNil tests that the provided value is not nil. If the value is nil, the assertion fails.
func NotNil(t *testing.T, value any) {
	t.Helper()
	v := reflect.ValueOf(value)
	if !v.IsValid() || v.Kind() == reflect.Pointer && reflect.ValueOf(value).IsNil() {
		t.Fatalf("Unexpected nil value")
	}
}
