package mckeeman

import "fmt"

type Singleton struct {
	CodePoint CodePoint
}

func NewSingleton(r rune) Singleton {
	return Singleton{CodePoint: CodePoint{Rune: r}}
}

func (s Singleton) Accept(_ Grammar, str string) int {
	return s.CodePoint.Accept(str)
}

func (s Singleton) Print() {
	fmt.Print("'")
	s.CodePoint.Print()
	fmt.Print("'")
}
