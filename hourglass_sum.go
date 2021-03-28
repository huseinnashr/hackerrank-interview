package hackerrank_interview

func hourglassSum(arr [][]int32) int32 {
	max := -int32(^uint32(0)>>1) - 1
	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			sum := arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
