package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	assert.Equal(t, float64(60), function.Times(10, 2, 3))
	assert.Equal(t, float64(600), function.Times(100, 2, 3))
	assert.Equal(t, float64(60), function.Times(10, 2, 3))
	assert.Equal(t, float64(600), function.Times(100, 2, 3))
}
