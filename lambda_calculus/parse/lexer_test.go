package lambda_calculus_parse

import (
	"fmt"
	"parsium/assert"
	"parsium/parse"
	"testing"
)

func process(parser *parse.Lexer, input string) parse.TreeData {
	var output parse.TreeData

	parser.Output = func(token parse.TreeData) {
		output = token
	}

	for _, c := range []byte(input) {
		parser.Process(c)
	}
	parser.Flush()

	return output
}

func TestLexer(t *testing.T) {
	t.Run("Process '\\'", func(t *testing.T) {
		l := NewLexer()
		l.Output = func(token parse.TreeData) {
			assert.Equal(t, token.Type, "\\")
		}
		l.Process('\\')
	})
	t.Run("Process '('", func(t *testing.T) {
		l := NewLexer()
		l.Output = func(token parse.TreeData) {
			assert.Equal(t, token.Type, "(")
		}
		l.Process('(')
	})
	t.Run("Output name", func(t *testing.T) {
		l := NewLexer()
		token := process(l, "abc")
		assert.Equal(t, token, parse.TreeData{
			Type:  "name",
			First: 0,
			Last:  3,
		})
	})
}

func TestLexerSandbox(t *testing.T) {
	lexer := NewLexer()
	lexer.Output = func(token parse.TreeData) {
		fmt.Println(token)
	}

	input := "\\f.\\a.(f a)"

	for _, c := range []byte(input) {
		lexer.Process(c)
	}
	lexer.Flush()
}
