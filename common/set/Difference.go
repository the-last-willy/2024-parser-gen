package set

func Difference[T comparable](a, b Set[T]) SimpleSet[T] {
	set := NewSimpleSet[T]()

	for e := range a.Elements() {
		if !b.Has(e) {
			set = set.Add(e).(SimpleSet[T])
		}
	}

	return set
}
