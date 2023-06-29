package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func main() {
	tree := Node{1, &Node{2, &Node{4, nil, nil}, &Node{5, nil, nil}}, &Node{3, &Node{6, nil, nil}, &Node{7, nil, nil}}}
	tree.PreorderTra()
}

func (root *Node) PreorderTra() {
	if root != nil {
		fmt.Printf("%d", root.Data)
		root.Left.PreorderTra()
		root.Right.PreorderTra()
	}

	return
}
