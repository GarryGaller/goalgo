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

func SelectionSort(arr []int) {
	/*
    Approach:
    - Repeatedly select the next smallest element from the unsorted array and
    move it to the front.
    
    Cost:
    - O(n^2) time and O(1) space.
    */
    
    var n = len(arr)
	for i := 0; i < n - 1; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

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
