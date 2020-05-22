package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestDivFloat(t *testing.T) {
	assert.Equal(t, float64(3), function.DivFloat(0.6, 0.2))
}
