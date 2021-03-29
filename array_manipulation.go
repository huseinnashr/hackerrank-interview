package hackerrank_interview

func arrayManipulation(n int32, queries [][]int32) int64 {
	array := make([]int64, n)
	maxValue := int64(0)
	for i := 0; i < len(queries); i++ {
		for j := queries[i][0] - 1; j < queries[i][1]; j++ {
			array[j] += int64(queries[i][2])
			if maxValue < array[j] {
				maxValue = array[j]
			}
		}
	}

	for i := 0; i < len(array); i++ {
		if maxValue < array[i] {
			maxValue = array[i]
		}
	}

	return maxValue
}
