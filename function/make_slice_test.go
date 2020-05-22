package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []interface{}{6, 2, "a", "b"}, function.MakeSlice(6, 2, "a", "b"))
}
