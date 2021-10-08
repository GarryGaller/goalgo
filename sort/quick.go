package sort

func partition(arr []int, low, high int) ([]int, int) {
	/*Partition Lomuto */
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func sort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = sort(arr, low, p-1)
		arr = sort(arr, p+1, high)
	}
	return arr
}

func QuickSort(arr []int) []int {
	/*
	   Approach:
	   - Recursively divide the input into two smaller arrays around a pivot, where
	     one half has items smaller than the pivot, other half has items bigger than
	     the pivot.
	   Cost:
	   - O(n*logn) time and O(n*logn) space.
	*/

	sort(arr, 0, len(arr)-1)

	return arr
}
