package sort

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
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
		} else {
			i++
		}
	}
}
