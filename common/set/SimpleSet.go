package set

import (
	"iter"
	"maps"
)

type SimpleSet[T comparable] struct {
	elements map[T]any
}

func NewSimpleSet[T comparable](elements ...T) SimpleSet[T] {
	ma := map[T]any{}
	for _, element := range elements {
		ma[element] = nil
	}
	return SimpleSet[T]{elements: ma}
}

// Builder interface

func (s SimpleSet[T]) Add(e T) Builder[T] {
	newElements := map[T]any{}
	for k, v := range s.elements {
		newElements[k] = v
	}
	newElements[e] = nil
	return SimpleSet[T]{elements: newElements}
}

func (s SimpleSet[T]) Remove(e T) Builder[T] {
	newElements := map[T]any{}
	for k, _ := range s.elements {
		if k != e {
			newElements[k] = nil
		}
	}
	return SimpleSet[T]{elements: newElements}
}

// Set interface

func (s SimpleSet[T]) Has(e T) bool {
	_, has := s.elements[e]
	return has
}

func (s SimpleSet[T]) Elements() iter.Seq[T] {
	return maps.Keys(s.elements)
}

func (s SimpleSet[T]) Len() int {
	return len(s.elements)
}
