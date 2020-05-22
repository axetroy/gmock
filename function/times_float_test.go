package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimesFloat(t *testing.T) {
	assert.Equal(t, float64(60), function.TimesFloat(10, 2, 3))
	assert.Equal(t, float64(600), function.TimesFloat(100, 2, 3))
}
