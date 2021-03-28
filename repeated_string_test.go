package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedString(t *testing.T) {
	assert.EqualValues(t, 7, repeatedString("aba", 10))
}
