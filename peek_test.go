package peek_test

import (
	"testing"

	peek "github.com/draganm/go-peek"
	"github.com/stretchr/testify/assert"
)

func TestStructures(t *testing.T) {

	cases := []struct {
		name          string
		path          string
		data          interface{}
		expectedValue interface{}
		expectedError error
	}{
		{
			name:          "simple_struct",
			path:          "Foo",
			data:          struct{ Foo string }{Foo: "bar"},
			expectedValue: "bar",
			expectedError: nil,
		},
		{
			name:          "simple_struct_pointer",
			path:          "Foo",
			data:          &struct{ Foo string }{Foo: "bar"},
			expectedValue: "bar",
			expectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)
			val, err := peek.Peek(c.path, c.data)
			assert.Equal(c.expectedError, err, "Unexpected error. Expected %v, got %v")
			assert.Equal(c.expectedValue, val, "Unexpected value. Expected %v, got %v")
		})
	}
}
