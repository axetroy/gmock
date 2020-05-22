package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimesInt(t *testing.T) {
	assert.Equal(t, int64(60), function.TimesInt(10, 2, 3))
	assert.Equal(t, int64(600), function.TimesInt(100, 2, 3))
}
