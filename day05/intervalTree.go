package main

import "fmt"


type IntervalTree struct {
	Root *Node
}

type Interval struct {
	Low    int
	High   int
	Offset int
}

type Node struct {
	Range *Interval
	Max   int
	Left  *Node
	Right *Node
}

func newNode(x Interval) *Node {
	node := Node{}
	node.Range = &x
	node.Max = x.High
	return &node
}

func (tree *IntervalTree) insert(x Interval) {
	tree.Root = insertNode(tree.Root, x)
}

func insertNode(root *Node, x Interval) *Node {
	if root == nil {
		return newNode(x)
	}
	if x.Low < root.Range.Low {
		root.Left = insertNode(root.Left, x)
	} else {
		root.Right = insertNode(root.Right, x)
	}
	if root.Max < x.High {
		root.Max = x.High
	}
	return root
}

func (tree *IntervalTree) search(x int) *Interval {
	return tree.Root.search(x)
}

func (root *Node) search(x int) *Interval {
	if root == nil {
		return &Interval{Offset: 0}
	}

	if x <= root.Range.High && x >= root.Range.Low {
		return root.Range
	}
	if root.Left != nil && root.Left.Max > x {
		return root.Left.search(x)
	}
	return root.Right.search(x)
}

func (root *Node) searchInterval(x *Interval) (res []*Interval) {
	if root == nil {
		return []*Interval{}
	}

	if root.Range.Low <= x.High && x.Low <= root.Range.High {
		res = append(res, root.Range)
	}
	if root.Left != nil && root.Left.Max > x.Low {
		res = append(res, root.Left.searchInterval(x)...)
	}
	res = append(res, root.Right.searchInterval(x)...)
	return
}

func (tree *IntervalTree) display() {
	tree.Root.display()
}

func (root *Node) display() {
	if root == nil {
		return
	}
	root.Left.display()
	fmt.Println("[", root.Range.Low, ", ", root.Range.High, "[ max = ", root.Max)
	root.Right.display()
}