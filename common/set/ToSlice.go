package set

func ToSlice[T any](s Set[T]) []T {
	sl := []T{}
	for e := range s.Elements() {
		sl = append(sl, e)
	}
	return sl
}
