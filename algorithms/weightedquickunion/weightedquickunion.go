package weightedquickunion

type WeightedQuickUnionFind interface {
	Root(i int) int
	Connected(p, q int)
	Union(p, q int)
}

type WeightedQuickUnionFindImpl struct {
	id   []int
	size []int
}

func NewWeightedQuickUnionFind(n int) *WeightedQuickUnionFindImpl {
	wg := &WeightedQuickUnionFindImpl{
		id:   make([]int, n),
		size: make([]int, n),
	}
	for i := 0; i < n; i++ {
		wg.id[i] = i
		wg.size[i] = 1
	}
	return wg
}

func (w *WeightedQuickUnionFindImpl) Root(i int) int {
	for i != w.id[i] {
		w.id[i] = w.id[w.id[i]] // path compression to halve path length
		i = w.id[i]
	}
	return i
}

func (w *WeightedQuickUnionFindImpl) Connected(p, q int) bool {
	return w.Root(p) == w.Root(q)
}

func (w *WeightedQuickUnionFindImpl) Union(p, q int) {
	i := w.Root(p)
	j := w.Root(q)
	if i == j {
		return
	}
	if w.size[i] < w.size[j] {
		w.id[i] = j
		w.size[j] += w.size[i]
	} else {
		w.id[j] = i
		w.size[i] += w.size[j]
	}
}
