package quickfind

var id []int

func QuickFind(n int) {
	id = make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
}

func union(p int, q int) {
	pId := id[p]
	qId := id[q]
	for i := 0; i < len(id); i++ {
		if id[i] == pId {
			id[i] = qId
		}
	}
}

func connected(p int, q int) bool {
	return id[p] == id[q]
}

func find(p int, q int) bool {
	return id[p] == id[q]
}
