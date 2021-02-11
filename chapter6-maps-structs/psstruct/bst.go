/*

insert
*/
package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (t *Tree) insert(val int) {
	if t.root == nil {
		t.root = &Node{val, nil, nil}
	} else {
		t.root.insert(val)
	}

}

func (n *Node) insert(val int) {
	if val <= n.val {
		if n.left == nil {
			n.left = &Node{val, nil, nil}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val, nil, nil}
		} else {
			n.right.insert(val)
		}
	}
}

/*
In Order Traversal
1- Traverse left subtree
2- visit root
3- Traverse right subtree
*/
func InOrder(n *Node) {
	if n == nil {
		return
	}
	InOrder(n.left)
	fmt.Println(n.val)
	InOrder(n.right)
}

/*
PreOrder Traversal
1- Visit root
2- left
3- right
*/
func PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	PreOrder(n.left)
	PreOrder(n.right)
}

/*
Post Traversal
1- left
2- right
3- root
*/
func PostOrder(n *Node) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	fmt.Println(n.val)
}

func main() {

	tree := new(Tree)
	tree.insert(3)
	tree.insert(2)
	tree.insert(4)

	InOrder(tree.root)
	fmt.Println()
	PreOrder(tree.root)
	fmt.Println()
	PostOrder(tree.root)
	fmt.Println()
}
