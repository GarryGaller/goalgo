package sort

type MaxHeap struct {
	slice []int
	size  int
}

func BuildMaxHeap(slice []int) MaxHeap {
	h := MaxHeap{slice: slice, size: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h MaxHeap) heapify(root int) {
	l, r := 2*root+1, 2*root+2
	max := root

	if l < h.size && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.size && h.slice[r] > h.slice[max] {
		max = r
	}

	if max != root {
		h.slice[root], h.slice[max] = h.slice[max], h.slice[root]
		h.heapify(max)
	}
}

func HeapSort(slice []int) []int {
	h := BuildMaxHeap(slice)

	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.size--
		h.heapify(0)
	}
	return h.slice
}
