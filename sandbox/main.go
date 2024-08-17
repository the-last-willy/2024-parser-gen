package main

import (
	"os"
	"parsium/format/json"
)

func main() {
	jsonf, err := os.ReadFile("/media/willy/Data/projects/parsium/assets/json/example-4.json")
	if err != nil {
		panic(err)
	}

	jsonp := json.NewParser()
	json := jsonp.Parse(string(jsonf))

	json.Root = json.Root.Derecursified()

	json.Root.PrintIndent(string(jsonf), 0)
}
