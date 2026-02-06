package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}
func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}
func (s *Stack[T]) Len() int {
	return len(s.elements)
}
func main() {
	eg1 := &Stack[int]{
		elements: []int{2, 2, 4, 5, 1},
	}
	eg2 := &Stack[string]{
		elements: []string{"li", "ru", "zh", "en"},
	}
	fmt.Println(eg1.elements, eg1.Len())
	fmt.Println(eg1.Pop())
	fmt.Println(eg1.elements, eg1.Len())
	eg1.Push(99)
	fmt.Println(eg1.elements, eg1.Len())
	fmt.Println(eg2.elements, eg2.Len())
	fmt.Println(eg2.Pop())
	fmt.Println(eg2.elements, eg2.Len())
	eg2.Push("99")
	fmt.Println(eg2.elements, eg2.Len())
}
