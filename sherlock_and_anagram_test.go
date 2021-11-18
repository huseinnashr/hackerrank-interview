package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSherlockAndAnagrams10(t *testing.T) {
	assert.EqualValues(t, 10, sherlockAndAnagrams("kkkk"))
}

func TestSherlockAndAnagrams3(t *testing.T) {
	assert.EqualValues(t, 3, sherlockAndAnagrams("ifailuhkqq"))
}
