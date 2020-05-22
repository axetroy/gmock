package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusInt(t *testing.T) {
	assert.Equal(t, int64(15), function.PlusInt(10, 2, 3))
	assert.Equal(t, int64(300), function.PlusInt(100+200))
}
