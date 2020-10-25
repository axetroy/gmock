package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	assert.Equal(t, float64(3), function.Div(6, 2))
	assert.Equal(t, float64(2.9999999999999996), function.Div(0.6, 0.2))
}
