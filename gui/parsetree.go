package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"parsium/parse"
	"parsium/tree"
)

type ParseTree struct {
	widget.Tree

	ParseTree tree.Tree[parse.TreeData]

	uidToNodeMap map[string]tree.Node
}

func NewParseTree() *ParseTree {
	pt := &ParseTree{}

	pt.ChildUIDs = pt.childUIDs
	pt.IsBranch = pt.isBranch
	pt.CreateNode = pt.createNode
	pt.UpdateNode = pt.updateNode

	pt.ExtendBaseWidget(pt)

	return pt
}

func (pt *ParseTree) SetTree(t tree.Tree[parse.TreeData]) {
	pt.ParseTree = t

	pt.uidToNodeMap = make(map[string]tree.Node)

	pt.uidToNodeMap[""] = *t.Root()

	tr := tree.NewTraverser[parse.TreeData]()
	tr.Enter = func(st tree.SubTree[parse.TreeData]) tree.TraverserCommand {
		pt.uidToNodeMap[pt.NodeToId(*st.Root())] = *st.Root()
		return tree.TraverserContinue
	}
	tr.Process(t)
}

func (pt *ParseTree) NodeToId(n tree.Node) string {
	return fmt.Sprint(n)
}

func (pt *ParseTree) IdToNode(uid string) tree.Node {
	node, found := pt.uidToNodeMap[uid]
	if !found {
		panic("uid not found")
	}
	return node
}

// widget.Tree callbacks

func (pt *ParseTree) childUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	if pt.ParseTree == nil {
		return []widget.TreeNodeID{}
	}

	// Root
	if id == "" {
		if pt.ParseTree.Root() == nil {
			return []widget.TreeNodeID{}
		}
		return []widget.TreeNodeID{pt.NodeToId(*pt.ParseTree.Root())}
	}

	// Node
	uids := []widget.TreeNodeID{}
	for _, child := range pt.ParseTree.ChildrenOf(pt.IdToNode(id)) {
		uids = append(uids, pt.NodeToId(child))
	}
	return uids
}

func (pt *ParseTree) isBranch(id widget.TreeNodeID) bool {
	if pt.ParseTree == nil || pt.ParseTree.Root() == nil {
		return false
	}

	node := pt.IdToNode(id)
	return len(pt.ParseTree.ChildrenOf(node)) > 0
}

func (pt *ParseTree) createNode(branch bool) fyne.CanvasObject {
	if branch {
		return widget.NewLabel("Branch template")
	}
	return widget.NewLabel("Leaf template")
}

func (pt *ParseTree) updateNode(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	node := pt.IdToNode(id)
	data := pt.ParseTree.DataOf(node)
	text := fmt.Sprintf("%s [%d, %d)", data.Type, data.First, data.Last)
	o.(*widget.Label).SetText(text)
}
