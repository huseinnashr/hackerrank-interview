package hackerrank_interview

/*
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func convertArrayToMap(arr []string) map[string]int {
	mappedArr := make(map[string]int)
	for _, el := range arr {
		mappedArr[el] += 1
	}
	return mappedArr
}

func checkMagazine(magazine []string, notes []string) string {
	mappedMagazine := convertArrayToMap(magazine)

	for _, note := range notes {
		count := mappedMagazine[note]

		if count == 0 {
			return "No"
		}

		mappedMagazine[note] -= 1
	}
	return "Yes"
}
