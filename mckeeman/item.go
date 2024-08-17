package mckeeman

type Item interface {
	Accept(Grammar, string) int
	Print()
}
