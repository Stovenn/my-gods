package quickunion

var id []int

func QuickUnion(n int) {
	id = make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
}

func root(i int) int {
	for i != id[i] {
		i = id[i]
	}
	return i
}

func connected(p, q int) bool {
	return root(p) == root(q)
}

func union(p, q int) {
	id[root(p)] = root(q)
}
