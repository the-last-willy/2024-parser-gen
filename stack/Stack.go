package stack

import "slices"

type Stack[T any] interface {
	At(int) T
	Len() int

	Push(T) Stack[T]
	Pop() Stack[T]
}

type SimpleStack[T any] struct {
	Slice []T
}

func NewStack[T any](elements ...T) Stack[T] {
	reversed := append([]T{}, elements...)
	slices.Reverse(reversed)
	return SimpleStack[T]{Slice: reversed}
}

// Stack interface

func (s SimpleStack[T]) At(i int) T {
	return s.Slice[len(s.Slice)-i-1]
}

func (s SimpleStack[T]) Len() int {
	return len(s.Slice)
}

func (s SimpleStack[T]) Pop() Stack[T] {
	res := s
	res.Slice = s.Slice[:s.Len()-1]
	return res
}

func (s SimpleStack[T]) Push(element T) Stack[T] {
	res := s
	res.Slice = append(s.Slice, element)
	return res
}
