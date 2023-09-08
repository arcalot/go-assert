package assert_test

import (
	"go.arcalot.io/assert"
	"testing"
)

//nolint:dupl
func TestMapContains(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	result := assert.MapContainsKey(t, "a", input)
	assert.Equals(t, result, 1)
	assert.MapNotContainsKey(t, "d", input)

	testFailure(t, func(t *testing.T) {
		assert.MapContainsKey(t, "z", input)
	})
	testFailure(t, func(t *testing.T) {
		assert.MapNotContainsKey(t, "a", input)
	})

	input2 := map[int]string{1: "a", 2: "b"}
	result2 := assert.MapContainsKey(t, 1, input2)
	assert.Equals(t, result2, "a")
	assert.MapNotContainsKey(t, 3, input2)

	testFailure(t, func(t *testing.T) {
		assert.MapContainsKey(t, 3, input2)
	})
	testFailure(t, func(t *testing.T) {
		assert.MapNotContainsKey(t, 1, input2)
	})
}

//nolint:dupl
func TestMapContainsAny(t *testing.T) {
	input := map[any]int{"a": 1, "b": 2}
	result := assert.MapContainsKeyAny(t, "a", input)
	assert.Equals(t, result, 1)
	assert.MapNotContainsKeyAny(t, "d", input)

	testFailure(t, func(t *testing.T) {
		assert.MapContainsKeyAny(t, "z", input)
	})
	testFailure(t, func(t *testing.T) {
		assert.MapNotContainsKeyAny(t, "a", input)
	})

	input2 := map[any]string{1: "a", 2: "b"}
	result2 := assert.MapContainsKeyAny(t, 1, input2)
	assert.Equals(t, result2, "a")
	assert.MapNotContainsKeyAny(t, 3, input2)

	testFailure(t, func(t *testing.T) {
		assert.MapContainsKeyAny(t, 3, input2)
	})
	testFailure(t, func(t *testing.T) {
		assert.MapNotContainsKeyAny(t, 1, input2)
	})
}
