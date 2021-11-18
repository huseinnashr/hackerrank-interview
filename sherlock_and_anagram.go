package hackerrank_interview

import (
	"sort"
	"strings"
)

func sherlockAndAnagrams(s string) int32 {
	subAnagramPerm := int32(0)
	for size := 1; size < len(s); size++ {
		currPermSize := int32(0)
		for currOffset := 0; currOffset < len(s)-size; currOffset++ {
			splittedCurrSubs := strings.Split(s[currOffset:currOffset+size], "")
			sort.Strings(splittedCurrSubs)
			currSubs := strings.Join(splittedCurrSubs, "")

			for offset := currOffset + 1; offset < len(s)-(size-1); offset++ {
				splittedSubs := strings.Split(s[offset:offset+size], "")
				sort.Strings(splittedSubs)
				subs := strings.Join(splittedSubs, "")

				if currSubs == subs {
					currPermSize++
				}
			}

		}

		subAnagramPerm += currPermSize
	}

	return subAnagramPerm
}
