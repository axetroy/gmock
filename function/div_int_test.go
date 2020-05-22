package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestDivInt(t *testing.T) {
	assert.Equal(t, int64(3), function.DivInt(6, 2))
}
