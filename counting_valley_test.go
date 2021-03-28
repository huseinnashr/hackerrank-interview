package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingValley(t *testing.T) {
	assert.EqualValues(t, 3, countingValley(8, "DDUUDDUUDDUU"))
}
