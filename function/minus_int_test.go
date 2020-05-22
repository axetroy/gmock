package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinusInt(t *testing.T) {
	assert.Equal(t, int64(5), function.MinusInt(10, 2, 3))
	assert.Equal(t, int64(1), function.MinusInt(3, 2))
}
