package mckeeman

type Alternatives struct {
	Alternatives []Alternative
}

func (a Alternatives) Accept(g Grammar, s string) int {
	for _, a := range a.Alternatives {
		c := a.Accept(g, s)
		if c >= 0 {
			return c
		}
	}
	return -1
}

func (a Alternatives) Print() {
	for _, a := range a.Alternatives {
		a.Print()
	}
}
