package mckeeman

type Literal interface {
	Accept(Grammar, string) int
	Print()
}
