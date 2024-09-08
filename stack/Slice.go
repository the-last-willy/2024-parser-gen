package stack

func Slice[T any](s Stack[T], first, last any) []T {
	lb := 0
	if i, ok := first.(int); ok {
		lb = i
	}

	ub := s.Len()
	if i, ok := last.(int); ok {
		ub = i
	}

	sl := []T{}
	for i := lb; i < ub; i++ {
		sl = append(sl, s.At(i))
	}

	return sl
}
