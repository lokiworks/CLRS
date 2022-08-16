package main

import "fmt"

type BinaryTreeNode struct {
	k     int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func inorderTreeWalk(n *BinaryTreeNode) {
	if n != nil {
		inorderTreeWalk(n.left)
		fmt.Println(n.k)
		inorderTreeWalk(n.right)
	}
}

func treeSearch(n *BinaryTreeNode, k int) *BinaryTreeNode {
	if n == nil || k == n.k {
		return n
	}
	if k < n.k {
		return treeSearch(n.left, k)
	}
	return treeSearch(n.right, k)
}

func iterativeTreeSearch(n *BinaryTreeNode, k int) *BinaryTreeNode {
	for n != nil && k != n.k {
		if k < n.k {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

func main() {

	left := BinaryTreeNode{5, &BinaryTreeNode{3, nil, nil}, &BinaryTreeNode{6, nil, nil}}

	right := BinaryTreeNode{10, &BinaryTreeNode{8, nil, nil}, &BinaryTreeNode{15, nil, nil}}

	k := 7

	root := BinaryTreeNode{k, &left, &right}
	inorderTreeWalk(&root)

	s1 := treeSearch(&root, 5)
	fmt.Println(s1.k)

	s2 := iterativeTreeSearch(&root, 15)
	fmt.Println(s2.k)

}
