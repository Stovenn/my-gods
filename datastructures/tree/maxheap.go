package tree

type MaxHeap struct {
	items []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{items: []int{}}
}

func (h *MaxHeap) Insert(i int) {
	h.items = append(h.items, i)
	h.maxHeapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.items[0]
	last := len(h.items) - 1
	if len(h.items) == 0 {
		return -1
	}

	h.items[0] = h.items[last]
	h.items = h.items[:last]
	//heapify down
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	if len(h.items) == 1 {
		return
	}
	for h.items[parent(index)] < h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.items) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.items[l] > h.items[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.items[index] < h.items[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
