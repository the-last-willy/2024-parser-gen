package parse

import (
	"fmt"
	"testing"
	"unicode"
)

func TestLexer(t *testing.T) {
	lexer := NewLexer()
	lexer.Grammar = func(buffer string) *TreeData {
		if unicode.IsDigit(rune(buffer[0])) {
			i := 1
			for ; i < len(buffer); i++ {
				if !unicode.IsDigit(rune(buffer[i])) {
					break
				}
			}
			return &TreeData{
				Type:  "int",
				First: 0,
				Last:  i,
			}
		}

		if unicode.IsLetter(rune(buffer[0])) {
			i := 1
			for ; i < len(buffer); i++ {
				if !unicode.IsLetter(rune(buffer[i])) {
					break
				}
			}
			return &TreeData{
				Type:  "id",
				First: 0,
				Last:  i,
			}
		}

		return &TreeData{
			Type:  string(buffer[0]),
			First: 0,
			Last:  1,
		}
	}
	lexer.Output = func(token TreeData) {
		fmt.Println(token)
	}

	input := "A*2+1 "

	for _, c := range []byte(input) {
		lexer.Process(c)
	}
}
