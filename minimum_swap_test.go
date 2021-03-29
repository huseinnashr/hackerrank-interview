package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwaps(t *testing.T) {
	assert.EqualValues(t, 5, minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}))
	assert.EqualValues(t, 3, minimumSwaps([]int32{4, 3, 1, 2}))
	assert.EqualValues(t, 3, minimumSwaps([]int32{2, 3, 4, 1, 5}))
}
