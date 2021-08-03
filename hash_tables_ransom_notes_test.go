package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRansomNotesYes(t *testing.T) {
	assert.EqualValues(t, "Yes", checkMagazine(
		[]string{"give", "me", "one", "grand", "today", "night"},
		[]string{"give", "one", "grand", "today"},
	))
}

func TestRansomNotesNo(t *testing.T) {
	assert.EqualValues(t, "No", checkMagazine(
		[]string{"two", "times", "three", "is", "not", "four"},
		[]string{"two", "times", "two", "is", "four"},
	))
}

func TestRansomNotesNo2(t *testing.T) {
	assert.EqualValues(t, "No", checkMagazine(
		[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
		[]string{"ive", "got", "some", "coconuts"},
	))
}
