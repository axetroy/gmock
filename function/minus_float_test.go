package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinusFloat(t *testing.T) {
	assert.Equal(t, float64(3), function.MinusFloat(6, 3))
	assert.Equal(t, 0.19999999999999998, function.MinusFloat(0.3, 0.1))
}
