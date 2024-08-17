package json_test

import (
	"os"
	"parsium/assert"
	"parsium/format/json"
	"parsium/parse"
	"testing"
)

func ParseFile(path string) parse.DataTree {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	p := json.NewParser()
	return p.Parse(string(f))
}

func TestExample1(t *testing.T) {
	res := ParseFile("data/example-1.json")
	assert.NotNil(t, res.Root)
}

func TestExample2(t *testing.T) {
	res := ParseFile("data/example-2.json")
	assert.NotNil(t, res.Root)
}

func TestExample3(t *testing.T) {
	res := ParseFile("data/example-3.json")
	assert.NotNil(t, res.Root)
}

func TestExample4(t *testing.T) {
	res := ParseFile("data/example-4.json")
	assert.NotNil(t, res.Root)
}

func TestExample5(t *testing.T) {
	res := ParseFile("data/example-5.json")
	assert.NotNil(t, res.Root)
}
