package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(i T) {
	s.items = append(s.items, i)
}

func (s *Stack[T]) Pop() T {
	var popped T
	if s.size() == 0 {
		fmt.Println("The stack is empty")
		return *new(T)
	} else if s.size() == 1 {
		popped = s.items[0]
		s.items = s.items[:0]
	} else {
		popped = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
	return popped
}

func (s *Stack[T]) Print() {
	for _, n := range s.items {
		fmt.Print(n)
	}
}

func (s *Stack[T]) size() int {
	return len(s.items)
}
