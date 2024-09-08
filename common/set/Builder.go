package set

type Builder[T any] interface {
	Add(T) Builder[T]
	Remove(T) Builder[T]
}
