package stack

func HasPrefix[T comparable](s Stack[T], prefix ...T) bool {
	if len(prefix) > s.Len() {
		return false
	}

	for i := range prefix {
		if prefix[len(prefix)-i-1] != s.At(i) {
			return false
		}
	}

	return true
}
