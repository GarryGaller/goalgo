package sort

func MergeSort(in []int) []int {
	/*
	   Problem:
	   - Implement merge sort.
	   Approach:
	   - Split the input in half, recursively sorts each half, then merge the
	     sorted halves back together.
	   Cost:
	   - O(nlogn) time and O(n) space.
	*/

	if len(in) <= 1 {
		return in
	}

	middle := len(in) / 2
	left := in[:middle]
	right := in[middle:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) (out []int) {

	out = make([]int, 0)

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}

	if i < len(left) {
		out = append(out, left[i:]...)
	}

	if j < len(right) {
		out = append(out, right[j:]...)
	}

	return
}
