package sort

func SelectionSort(arr []int) {
	/*
	   Approach:
	   - Repeatedly select the next smallest element from the unsorted array and
	   move it to the front.

	   Cost:
	   - O(n^2) time and O(1) space.
	*/

	var n = len(arr)
	for i := 0; i < n-1; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
