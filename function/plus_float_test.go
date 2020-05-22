package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusFloat(t *testing.T) {
	assert.Equal(t, float64(15), function.PlusFloat(10, 2, 3))
	assert.Equal(t, 0.30000000000000004, function.PlusFloat(0.1, 0.2))
}
