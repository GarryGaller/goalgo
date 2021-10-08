package sort

// Swap two array values given their indices.
func Swap(v interface{}, i, j int) {
	switch a := v.(type) {
	case []int:
		SwapInt(a, i, j)
	case []string:
		SwapString(a, i, j)
	}
}

// SwapInt two array values given their indices.
func SwapInt(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// SwapString two array values given their indices.
func SwapString(a []string, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func СyclicSort(nums []int) {
	/*
	   Given an array containing n objects where each object, when created,
	   was assigned a unique number from 1 to n based on their creation sequence.

	   Approach:
	   - Use the fact that we are given a range of 1-n, can try placing each number at
	     its correct index in the array, say 1 at 0, 2 at 1, 3 at 2 and so on.
	   - Iterate through the array and if the current number is not at the correct index,
	     swap it with the number at its correct index.

	   Cost:
	   - O(n) time, O(1) space.

	*/

	i := 0
	for i < len(nums) {
		// if the current number is not at the correct index, swap it with the
		// one that is at the correct index.
		correctIndex := nums[i] - 1
		if nums[i] != nums[correctIndex] {
			Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}
}

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
