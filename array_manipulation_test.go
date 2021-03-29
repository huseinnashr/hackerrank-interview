package hackerrank_interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayManipulation(t *testing.T) {
	assert.EqualValues(t, 10, arrayManipulation(10, [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}))
	assert.EqualValues(t, 200, arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
}

func BenchmarkArrayManipulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayManipulation(1000, [][]int32{{1, 500, 3}, {40, 80, 7}, {6, 900, 1}, {10, 50, 3}, {400, 800, 7}, {600, 900, 1}})
	}
}
