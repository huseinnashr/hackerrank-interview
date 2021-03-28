package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpingOnCloud(t *testing.T) {
	assert.EqualValues(t, 3, jumpingOnCloud([]int32{0, 1, 0, 0, 0, 1, 0}))
}
