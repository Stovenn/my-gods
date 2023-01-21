package tree

type Heap interface {
	Insert(i int)
	Extract() int
	swap(i, j int)
}

func left(parentIndex int) int {
	return parentIndex*2 + 1
}

func right(parentIndex int) int {
	return parentIndex*2 + 2
}

func parent(childIndex int) int {
	return (childIndex - 1) / 2
}
