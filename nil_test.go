package assert_test

import (
	"testing"
	"unsafe"

	"go.arcalot.io/assert"
)

type TestInterface interface {
}

func TestNil(t *testing.T) {
	type testStruct struct{}
	var pointerVal *testStruct = nil
	var mapVal map[any]any = nil
	var chanVal chan int = nil
	var funcVal func() int = nil
	var interfaceVal TestInterface = nil
	var sliceVal []int = nil
	var unsafePointerVal unsafe.Pointer = nil
	assert.Nil(t, nil)
	assert.Nil(t, pointerVal)
	assert.Nil(t, mapVal)
	assert.Nil(t, chanVal)
	assert.Nil(t, funcVal)
	assert.Nil(t, interfaceVal)
	assert.Nil(t, sliceVal)
	assert.Nil(t, unsafePointerVal)

	testFailure(t, func(t *testing.T) {
		assert.Nil(t, "") // Non-pointer literal
	})
	testFailure(t, func(t *testing.T) {
		assert.Nil(t, &testStruct{}) // Pointer
	})
	testFailure(t, func(t *testing.T) {
		assert.Nil(t, make(map[any]any)) // map
	})
	testFailure(t, func(t *testing.T) {
		assert.Nil(t, make(chan int)) // Channel
	})
	testFailure(t, func(t *testing.T) {
		assert.Nil(t, func() {}) // function
	})
	testFailure(t, func(t *testing.T) {
		var test TestInterface = "test"
		assert.Nil(t, test) // interface
	})
	testFailure(t, func(t *testing.T) {
		assert.Nil(t, make([]int, 0)) // slice
	})
}

func TestNotNil(t *testing.T) {
	type testStruct struct{}
	var v = &testStruct{}
	assert.NotNil(t, v)
	assert.NotNil(t, "Hello world!")
	testFailure(t, func(t *testing.T) {
		assert.NotNil(t, nil)
	})
	testFailure(t, func(t *testing.T) {
		var mapVal map[any]any = nil
		assert.NotNil(t, mapVal)
	})
}
