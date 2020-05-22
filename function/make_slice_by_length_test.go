package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSliceByLength(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, function.MakeSliceByLength(6))
	assert.Equal(t, []int{0, 1, 2}, function.MakeSliceByLength(3))
}
