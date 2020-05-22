package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRandomStr(t *testing.T) {
	assert.Len(t, function.RandomStr(6), 6)
	assert.Len(t, function.RandomStr(3), 3)
	assert.Len(t, function.RandomStr(32), 32)

	var str = function.RandomStr(6, "0123456789")
	assert.Len(t, str, 6)
	assert.True(t, regexp.MustCompile("^\\d{6}$").MatchString(str))
}
