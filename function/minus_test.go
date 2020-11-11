package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestMinusInt(t *testing.T) {
	assert.Equal(t, float64(5), function.Minus(10, 2, 3))
	assert.Equal(t, float64(1), function.Minus(3, 2))
	assert.Equal(t, float64(3), function.Minus(6, 3))
	assert.Equal(t, 0.19999999999999998, function.Minus(0.3, 0.1))
	assert.Equal(t, 0.3, function.Minus("0.5", 0.2))
}
