package tree

// The Builder interface does not have a Build function.
// The Build function should be provided
type Builder[Data any] interface {
	Tree(d Data, children []Builder[Data]) Builder[Data]

	For(Tree[Data]) Builder[Data]

	ConvertTo(Builder[Data]) Builder[Data]
}
