package common

func SliceLast[T any](s []T) T {
	return s[len(s)-1]
}
