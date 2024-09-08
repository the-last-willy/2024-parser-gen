package lambda_calculus_parse_test

import (
	"fmt"
	"parsium/assert"
	parse2 "parsium/lambda_calculus/parse"
	"parsium/parse"
	"testing"
)

func parserProcess(p *parse.BottomUpParser, tokens ...string) []parse.TreeData {
	var result []parse.TreeData
	p.Output = func(data parse.TreeData) {
		result = append(result, data)
	}

	for _, token := range tokens {
		p.Process(token)
	}
	p.Flush()

	return result
}

func TestParser(t *testing.T) {
	t.Run("Output expr for app", func(t *testing.T) {
		parser := parse2.NewParser()
		results := parserProcess(parser, "app")
		assert.SliceContain(t, results, parse.TreeData{
			Type:  "expr",
			First: 0,
			Last:  1,
		})
	})
	t.Run("Output expr for fun", func(t *testing.T) {
		parser := parse2.NewParser()
		results := parserProcess(parser, "fun")
		assert.SliceContain(t, results, parse.TreeData{
			Type:  "expr",
			First: 0,
			Last:  1,
		})
	})
	t.Run("Output expr for name", func(t *testing.T) {
		parser := parse2.NewParser()
		results := parserProcess(parser, "name")
		assert.SliceContain(t, results, parse.TreeData{
			Type:  "expr",
			First: 0,
			Last:  1,
		})
	})
	t.Run("Output app", func(t *testing.T) {
		parser := parse2.NewParser()
		results := parserProcess(parser, "(", "expr", "expr", ")")
		assert.SliceContain(t, results, parse.TreeData{
			Type:  "app",
			First: 0,
			Last:  4,
		})
	})
	t.Run("Output fun", func(t *testing.T) {
		parser := parse2.NewParser()
		results := parserProcess(parser, "\\", "name", ".", "expr")
		assert.SliceContain(t, results, parse.TreeData{
			Type:  "fun",
			First: 0,
			Last:  4,
		})
	})

	t.Run("Output fun 2", func(t *testing.T) {
		parser := parse2.NewParser()
		results := parserProcess(parser, "\\", "name", ".", "name")
		fmt.Print(parser.Stack)
		assert.SliceContain(t, results, parse.TreeData{
			Type:  "fun",
			First: 0,
			Last:  4,
		})
	})
}
