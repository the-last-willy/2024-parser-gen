package set

import "iter"

type Set[T any] interface {
	Elements() iter.Seq[T]
	Has(T) bool
	Len() int
}
