package hackerrank_interview

func minimumSwaps(arr []int32) int32 {
	numOfSwap := int32(0)
	smallestNumber := 1

	for i := 0; i < len(arr); i++ {
		matchIndex := int32(i + smallestNumber)
		if arr[i] == matchIndex {
			continue
		}

		for j := i; j < len(arr); j++ {
			if arr[j] == matchIndex {
				arr[i], arr[j] = arr[j], arr[i]
				numOfSwap++
				break
			}
		}
	}

	return numOfSwap
}
