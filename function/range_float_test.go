package function_test

import (
	"testing"

	"github.com/axetroy/gmock/function"
	"github.com/stretchr/testify/assert"
)

func TestRangeFloat(t *testing.T) {

	{
		slice := make([]int, 1000)

		for _, _ = range slice {
			r := function.RangeFloat(1, 10)
			assert.True(t, r <= 10)
			assert.True(t, r >= 1)
		}
	}

}
