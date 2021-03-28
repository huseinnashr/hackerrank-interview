package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotLeft(t *testing.T) {
	assert.EqualValues(t, []int32{5, 1, 2, 3, 4}, rotLeft([]int32{1, 2, 3, 4, 5}, 4))
}
