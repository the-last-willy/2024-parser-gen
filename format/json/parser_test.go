package json_test

import (
	"os"
	"parsium/assert"
	"parsium/format/json"
	"parsium/parse"
	"parsium/tree"
	"testing"
)

func ParseFile(path string) tree.Tree[parse.TreeData] {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	p := json.NewParser()
	return p.Parse(string(f))
}

func TestExample1(t *testing.T) {
	res := ParseFile("data/example-1.json")
	assert.False(t, res.IsEmpty())
}

func TestExample2(t *testing.T) {
	res := ParseFile("data/example-2.json")
	assert.False(t, res.IsEmpty())
}

func TestExample3(t *testing.T) {
	res := ParseFile("data/example-3.json")
	assert.False(t, res.IsEmpty())
}

func TestExample4(t *testing.T) {
	res := ParseFile("data/example-4.json")
	assert.False(t, res.IsEmpty())
}

func TestExample5(t *testing.T) {
	res := ParseFile("data/example-5.json")
	assert.False(t, res.IsEmpty())
}
