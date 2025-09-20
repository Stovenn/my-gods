package usescases

import (
	"fmt"
	"my-gods/algorithms/weightedquickunion"
)

const (
	numberOfVirtualIndexes = 2
	indexOffset            = 1
)

type Percolation struct {
	grid               [][]bool
	gridWidth          int
	virtualBottomIndex int
	openSitesCounter   int
	uf                 *weightedquickunion.WeightedQuickUnionFindImpl
}

func NewPercolation(N int) (*Percolation, error) {
	p := &Percolation{}
	if N <= 0 {
		return nil, fmt.Errorf("n must be superior to 0")
	}
	p.gridWidth = N
	p.virtualBottomIndex = N*N + 1
	//grid initialization
	p.grid = make([][]bool, N)
	for i := 0; i < N; i++ {
		p.grid[i] = make([]bool, N)
	}
	p.uf = weightedquickunion.NewWeightedQuickUnionFind(N*N + numberOfVirtualIndexes)

	return p, nil
}

func (p *Percolation) open(row, col int) error {
	isOpen, err := p.isOpen(row, col)
	if err != nil {
		return err
	}
	if isOpen {
		return nil
	} else {
		p.grid[row-indexOffset][col-indexOffset] = true
		p.openSitesCounter++
		coords := p.convertTo1D(row, col)

		p.connectAdjacent(row-1, col, coords)
		p.connectAdjacent(row+1, col, coords)
		p.connectAdjacent(row, col-1, coords)
		p.connectAdjacent(row, col+1, coords)

		if row == 1 {
			p.uf.Union(0, coords)
		}
		if row == p.gridWidth {
			p.uf.Union(p.virtualBottomIndex, coords)
		}
		return nil
	}
}

func (p *Percolation) isOpen(row, col int) (bool, error) {
	if row-indexOffset < 0 || col-indexOffset < 0 {
		return false, fmt.Errorf("boundary error: row: %d, col: %d", row, col)
	}
	fmt.Println(row, col)
	return p.grid[row-indexOffset][col-indexOffset] == true, nil
}

func (p *Percolation) connectAdjacent(row, col, coords int) {
	o, err := p.isOpen(row, col)
	if err != nil {
	}
	if o {
		p.uf.Union(p.convertTo1D(row, col), coords)
	}
}

func (p *Percolation) convertTo1D(i, j int) int {
	return p.gridWidth*(i-1) + j
}

func (p *Percolation) isFull() bool {
	return false
}

func (p *Percolation) percolates() bool {
	return p.uf.Root(0) == p.uf.Root(p.virtualBottomIndex)
}
