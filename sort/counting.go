package sort

func СountingSort(in []int, max int) []int {
	/*
	   Applies to a set of numbers from a relatively small range

	   Cost:
	   - O(n) time and O(n) space.
	*/

	// utilize max value to create a fix-sized list of item counts.
	counts := make([]int, max+1)
	out := make([]int, 0)

	// populate the array where its indices represent items themselves and
	// its values represent how many time the item appears.
	for _, item := range in {
		counts[item]++
	}

	// iterate through the counts and add the item to the output list.
	for i := 0; i < len(counts); i++ {
		count := counts[i]

		for j := 0; j < count; j++ {
			out = append(out, i)
		}
	}

	return out
}
