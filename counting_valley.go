package hackerrank_interview

func countingValley(steps int32, path string) int32 {
	currLevel := 0
	isAtValley := false
	countAtValley := 0
	for _, c := range path {
		if c == 'U' {
			currLevel++
		} else {
			currLevel--
		}

		if !isAtValley && currLevel < 0 {
			countAtValley++
			isAtValley = true
		} else if currLevel >= 0 {
			isAtValley = false
		}
	}

	return int32(countAtValley)
}
