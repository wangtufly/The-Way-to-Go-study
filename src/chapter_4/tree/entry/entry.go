package main

import (
	"chapter_4/tree"
)

type myTreeNode struct {
	node *tree.Node
}

//后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	//
	//nodes := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//root.Print()
	//root.Right.Left.SetValue(4)
	//root.Right.Left.Print()
	//fmt.Println()

	root.Traverse()
	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	//fmt.Println()
}
