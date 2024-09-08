package common

func SliceReduceRange[T any](s []T, first, last int, r func([]T) []T) []T {
	prefix := s[:first]
	middle := s[first:last]
	suffix := s[last:]

	reduced := r(middle)

	return append(append(prefix, reduced...), suffix...)
}
