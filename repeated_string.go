package hackerrank_interview

import (
	"strings"
)

func repeatedString(s string, n int64) int64 {
	stringLen := int64(len(s))
	mod := n % stringLen
	stringMod := s[:mod]

	countStringA := int64(strings.Count(s, "a"))
	countStringModA := int64(strings.Count(stringMod, "a"))
	repeat := n / stringLen

	return countStringA*repeat + countStringModA
}
