package lambda_calculus

import (
	"parsium/parse"
	"slices"
	"unicode"
)

func NewLexer() *parse.Lexer {
	lexer := parse.NewLexer()
	lexer.Grammar = func(buffer string) *parse.TreeData {
		keySymbols := []uint8{'\\', '(', ')', '.'}

		if slices.Contains(keySymbols, buffer[0]) {
			return &parse.TreeData{
				Type:  string(buffer[0]),
				First: 0,
				Last:  1,
			}
		}

		if unicode.IsLetter(rune(buffer[0])) {
			i := 1
			for ; i < len(buffer); i++ {
				if !unicode.IsLetter(rune(buffer[i])) {
					return &parse.TreeData{
						Type:  "name",
						First: 0,
						Last:  i,
					}
				}
			}
			return nil
		}

		if unicode.IsSpace(rune(buffer[0])) {
			i := 1
			for ; i < len(buffer); i++ {
				if !unicode.IsSpace(rune(buffer[i])) {
					return &parse.TreeData{
						Type:  "space",
						First: 0,
						Last:  i,
					}
				}
			}
			return nil
		}

		return nil
	}
	return lexer
}
