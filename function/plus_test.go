package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	assert.Equal(t, float64(15), function.Plus(10, 2, 3))
	assert.Equal(t, float64(300), function.Plus(100+200))
	assert.Equal(t, float64(15), function.Plus(10, 2, 3))
	assert.Equal(t, 0.30000000000000004, function.Plus(0.1, 0.2))
}
