package hackerrank_interview

func jumpingOnCloud(c []int32) int32 {
	countJump := 0
	n := len(c)
	justDash := false
	justJump := false
	for i := 0; i < n-1; i++ {
		if c[i+1] == 0 && !justDash && justJump {
			justDash = true
			justJump = false
			continue
		}

		countJump++
		justJump = true
		justDash = false
	}

	return int32(countJump)
}
