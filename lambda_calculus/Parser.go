package lambda_calculus

import (
	"parsium/parse"
	"parsium/stack"
)

func NewParser() *parse.BottomUpParser {
	parser := parse.NewBottomUpParser()
	parser.Grammar = func(s stack.Stack[parse.Token], done bool) *parse.TreeData {
		if done {
			s = s.Push("")
		}

		d := func() *parse.TreeData {
			if stack.HasPrefix(s.Pop(), "app") {
				return &parse.TreeData{
					Type:  "expr",
					First: 1,
					Last:  2,
				}
			}

			if stack.HasPrefix(s.Pop(), "fun") {
				return &parse.TreeData{
					Type:  "expr",
					First: 1,
					Last:  2,
				}
			}

			if stack.HasPrefix(s.Pop(), "name") && s.At(0) != "." {
				return &parse.TreeData{
					Type:  "expr",
					First: 1,
					Last:  2,
				}
			}

			if stack.HasPrefix(s.Pop(), "(", "expr", "expr", ")") {
				return &parse.TreeData{
					Type:  "app",
					First: 1,
					Last:  5,
				}
			}

			if stack.HasPrefix(s.Pop(), `\`, "name", ".", "expr") {
				return &parse.TreeData{
					Type:  "fun",
					First: 1,
					Last:  5,
				}
			}

			return nil
		}()

		if done && d != nil {
			d.First -= 1
			d.Last -= 1
		}

		return d
	}
	return parser
}
