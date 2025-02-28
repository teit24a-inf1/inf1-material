package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {
	firstposition := -1
	lastposition := -1
	for position := range list {
		if list[position] == m {
			firstposition = position
		}

		if list[position] == n {
			lastposition = position
		}
	}

	if lastposition <= firstposition {
		return []int{}
	}

	result := []int{}
	nums := list[firstposition+1 : lastposition]
	//for pos := range result {
	for pos := range nums {
		//result[pos] = nums[pos]
		result = append(result, nums[pos])
	}

	return result
}
