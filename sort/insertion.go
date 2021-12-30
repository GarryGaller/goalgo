package sort

func InsertionSort(arr []int) {
	/*
	    Approach:
	   - Insert elements from an unsorted array into a sorted subsection of the
	     array, one item at a time.
	   Cost:
	   - O(n^2) time and O(1) space.
	*/

	var n = len(arr)
	var i int
	for i = 1; i < n; i++ {
		var j int
		j = i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}
