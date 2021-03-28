package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourglassSumPositive(t *testing.T) {
	assert.EqualValues(t, 19, hourglassSum([][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}))
}

func TestHourglassSumContainNegative(t *testing.T) {
	assert.EqualValues(t, 28, hourglassSum([][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}))
}

func TestHourglassSumNegative(t *testing.T) {
	assert.EqualValues(t, -7, hourglassSum([][]int32{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}))
}
