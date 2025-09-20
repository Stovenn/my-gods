package heap

type MinHeap []int

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) siftDown(currIdx int, endIdx int) {
	leftChild := currIdx*2 + 1
	for leftChild <= endIdx {
		rightChild := currIdx*2 + 2
		if rightChild > endIdx {
			rightChild = -1
		}

		targetIdx := leftChild
		if rightChild != -1 && (*h)[rightChild] < (*h)[leftChild] {
			targetIdx = rightChild
		}

		if (*h)[currIdx] > (*h)[targetIdx] {
			h.swap(currIdx, targetIdx)
			currIdx = targetIdx
			leftChild = currIdx*2 + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	currIdx := len(*h) - 1
	parentIdx := (currIdx - 1) / 2

	for currIdx > 0 && (*h)[currIdx] < (*h)[parentIdx] {
		h.swap(currIdx, parentIdx)
		currIdx = parentIdx
		parentIdx = (currIdx - 1) / 2
	}
}

func (h *MinHeap) BuildHeap() {
	startIdx := (len(*h) - 2) / 2
	for i := startIdx; i >= 0; i-- {
		endIdx := len(*h) - 1
		h.siftDown(i, endIdx)
	}
}

func (h *MinHeap) Remove() int {
	heapLen := len(*h)
	//swap root and last leaf node
	h.swap(0, heapLen-1)
	val := (*h)[heapLen-1]

	(*h) = (*h)[:heapLen-1]
	// call sift down
	h.siftDown(0, heapLen-2)

	return val
}

func (h *MinHeap) Insert(val int) {
	// insert at the end of the heap
	*h = append(*h, val)
	// call sift up
	h.siftUp()
}
