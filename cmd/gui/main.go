package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"parsium/format/mckeeman"
	"parsium/gui"
)

func main() {
	a := app.New()

	tree := gui.NewParseTree()
	tree.SetTree(mckeeman.GrammarTree())

	text := gui.NewParseSource()
	text.SetText(mckeeman.GrammarSource)

	tree.OnSelected = func(uid widget.TreeNodeID) {
		node := tree.IdToNode(uid)
		data := tree.ParseTree.DataOf(node)
		text.Highlight(data.First, data.Last)
		text.Refresh()
	}

	layout := container.NewGridWithColumns(2, tree, container.NewScroll(text))

	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(1280, 720))

	w.SetContent(layout)

	w.ShowAndRun()
}
