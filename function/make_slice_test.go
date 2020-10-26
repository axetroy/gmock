package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []interface{}{6, 2, "a", "b"}, function.MakeSlice(6, 2, "a", "b"))
}
