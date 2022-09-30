package assert_test

import (
	"testing"

	"go.arcalot.io/assert"
)

func TestNil(t *testing.T) {
	type testStruct struct{}
	var v *testStruct
	assert.Nil(t, nil)
	assert.Nil(t, v)

	testFailure(t, func(t *testing.T) {
		assert.Nil(t, "")
	})
	testFailure(t, func(t *testing.T) {
		assert.Nil(t, &testStruct{})
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
}
