package function_test

import (
	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeInt(t *testing.T) {

	{
		slice := make([]int, 1000)

		for _, _ = range slice {
			r := function.RangeInt(1, 10)
			assert.True(t, r <= 10)
			assert.True(t, r >= 1)
		}
	}

}
