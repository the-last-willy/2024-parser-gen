package parse

import (
	"fmt"
	"slices"
	"testing"
)

func process(p *BottomUpParser, symbol string) {
	fmt.Println(append(p.Stack, symbol))
	p.Process(symbol)
}

func TestBottomUpParser(t *testing.T) {
	parser := NewBottomUpParser()
	parser.Grammar = func(stack []string) *TreeData {
		stack = append([]string{}, stack...)
		slices.Reverse(stack)

		if stack[0] == "id" || stack[0] == "int" {
			return &TreeData{
				Type:  "Value",
				First: 0,
				Last:  1,
			}
		}

		if stack[1] != "*" && stack[0] == "Value" {
			return &TreeData{
				Type:  "Products",
				First: 0,
				Last:  1,
			}
		}

		if stack[2] == "Products" && stack[1] == "*" && stack[0] == "Value" {
			return &TreeData{
				Type:  "Products",
				First: 0,
				Last:  3,
			}
		}

		if stack[2] != "+" && stack[1] == "Products" && stack[0] != "*" {
			return &TreeData{
				Type:  "Sums",
				First: 1,
				Last:  2,
			}
		}

		if stack[2] == "Sums" && stack[1] == "+" && stack[0] == "Products" {
			return &TreeData{
				Type:  "Sums",
				First: 0,
				Last:  3,
			}
		}

		return nil
	}
	parser.Output = func(data TreeData) {
		fmt.Println("-", parser.Stack)
	}

	process(parser, "id")
	process(parser, "*")
	process(parser, "int")
	process(parser, "+")
	process(parser, "int")
}
