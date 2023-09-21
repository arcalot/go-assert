package assert_test

import (
	"go.arcalot.io/assert"
	"testing"
)

func TestSliceContainsInt(t *testing.T) {
	input := []int{1, 2}
	result := assert.SliceContains(t, 1, input)
	assert.Equals(t, result, 0)
	result = assert.SliceContains(t, 2, input)
	assert.Equals(t, result, 1)

	assert.SliceNotContains(t, 0, input)
	assert.SliceNotContains(t, -10, input)
	assert.SliceNotContains(t, 100, input)
	assert.SliceNotContains(t, "a", input)

	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, 0, input)
	})
	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, "a", input)
	})
	testFailure(t, func(t *testing.T) {
		assert.SliceNotContains(t, 1, input)
	})
}

func TestSliceContainsStr(t *testing.T) {
	input := []string{"a", "b"}
	result := assert.SliceContains(t, "a", input)
	assert.Equals(t, result, 0)
	result = assert.SliceContains(t, "b", input)
	assert.Equals(t, result, 1)

	assert.SliceNotContains(t, 0, input)
	assert.SliceNotContains(t, -10, input)
	assert.SliceNotContains(t, 100, input)
	assert.SliceNotContains(t, "", input)
	assert.SliceNotContains(t, "c", input)

	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, "c", input)
	})
	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, 0, input)
	})
	testFailure(t, func(t *testing.T) {
		assert.SliceNotContains(t, "a", input)
	})
}

func TestSliceContainsAnyEmpty(t *testing.T) {
	var emptyInput []any

	assert.SliceNotContains(t, 0, emptyInput)
	assert.SliceNotContains(t, -10, emptyInput)
	assert.SliceNotContains(t, 100, emptyInput)
	assert.SliceNotContains(t, "", emptyInput)
	assert.SliceNotContains(t, "c", emptyInput)

	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, "c", emptyInput)
	})
	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, 0, emptyInput)
	})
}

func TestSliceContainsAny(t *testing.T) {
	input := []any{1, "a"}
	result := assert.SliceContains(t, 1, input)
	assert.Equals(t, result, 0)
	result = assert.SliceContains(t, "a", input)
	assert.Equals(t, result, 1)

	assert.SliceNotContains(t, "b", input)
	assert.SliceNotContains(t, 0, input)

	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, 0, input)
	})
	testFailure(t, func(t *testing.T) {
		assert.SliceContains(t, "b", input)
	})
}

type sliceTestObject struct {
	a int
	b string
}

func TestSliceContainsMatch(t *testing.T) {
	testObj1 := sliceTestObject{1, "test1"}
	testObj2 := sliceTestObject{2, "test2"}
	testSlice := []sliceTestObject{testObj1, testObj2}

	assert.SliceContainsMatch(t, func(value sliceTestObject) bool {
		return value.a == 1
	}, testSlice)
	assert.SliceContainsMatch(t, func(value sliceTestObject) bool {
		return value.a == 2
	}, testSlice)
	assert.SliceContainsMatch(t, func(value sliceTestObject) bool {
		return value.b == "test1"
	}, testSlice)

	testFailure(t, func(t *testing.T) {
		assert.SliceContainsMatch(t, func(value sliceTestObject) bool {
			return value.a == 0
		}, testSlice)
	})
	testFailure(t, func(t *testing.T) {
		// Test empty slice
		assert.SliceContainsMatch(t, func(value sliceTestObject) bool {
			return value.a == 1
		}, []sliceTestObject{})
	})
}

func TestSliceContainsExtractor(t *testing.T) {
	testObj1 := sliceTestObject{1, "test1"}
	testObj2 := sliceTestObject{2, "test2"}
	testSlice := []sliceTestObject{testObj1, testObj2}

	assert.SliceContainsExtractor(t, func(value sliceTestObject) int {
		return value.a
	}, 1, testSlice)
	assert.SliceContainsExtractor(t, func(value sliceTestObject) int {
		return value.a
	}, 2, testSlice)
	assert.SliceContainsExtractor(t, func(value sliceTestObject) string {
		return value.b
	}, "test1", testSlice)

	testFailure(t, func(t *testing.T) {
		assert.SliceContainsExtractor(t, func(value sliceTestObject) int {
			return value.a
		}, 0, testSlice)
	})
	testFailure(t, func(t *testing.T) {
		// Test empty slice
		assert.SliceContainsExtractor(t, func(value sliceTestObject) int {
			return value.a
		}, 1, []sliceTestObject{})
	})
}
