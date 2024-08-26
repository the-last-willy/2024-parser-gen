package main

import (
	"fmt"
	"os"
	"parsium/format/mckeeman"
	"parsium/parse"
	"parsium/tree"
)

func main() {
	src, _ := os.ReadFile("/media/willy/Data/projects/parsium/assets/mckeeman.txt")
	//
	//p := mckeeman.NewParser()
	//res := p.Parse(string(src))
	//
	//res.Root = res.Root.Derecursified()
	//res = mckeeman.Simplify(res)
	//
	//tr := tree.NewSimpleTree[parse.TreeData]().AsTree()
	//tr = tr.WithRoot(res.Root.ToTreeNode(tr))
	//f := parse.TreeFormatter{}
	//file, _ := os.Create("tree.txt")
	//file.WriteString(f.Format(tr, string(src)))

	p2 := mckeeman.NewParser()
	b2 := p2.Parse(string(src), tree.NewSimpleTree[parse.TreeData]()).(tree.SimpleTree[parse.TreeData])

	tr2 := b2.Build()

	pd := mckeeman.NewSimplifier()
	tr2 = *pd.Process(tr2).(*tree.SimpleTree[parse.TreeData])

	pf := parse.TreeFormatter{}
	fmt.Println(pf.Format(tree.NewSubTree(tr2, tr2.Root()), string(src)))

	//res = mckeeman.Simplify(res)

	//nw := tree.NewSimpleTree[parse.TreeData]().AsTree()
	//nw = nw.WithRoot(res.Root.ToTreeNode(nw))

	//pf := parse.TreeFormatter{}
	//fmt.Println(pf.Format(nw, string(src)))

	//file, _ := os.Create("out.txt")
	//
	//gfmt := tree.GoFormatter[parse.TreeData]{}
	//gfmt.FormatData = func(i parse.TreeData) string {
	//	return fmt.Sprintf(`TreeData{
	// Type:  "%s",
	// First: %d,
	// Last:  %d,
	//}`,
	//		i.Type, i.First, i.Last)
	//}
	//_, _ = file.WriteString(gfmt.Format(tr))
}
