package main

import "fmt"

type BinaryTreeNode struct {
	k     int
	left  *BinaryTreeNode
	right *BinaryTreeNode
	p     *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func treeWalk(n *BinaryTreeNode) {
	if n != nil {
		treeWalk(n.left)
		fmt.Println(n.k)
		treeWalk(n.right)
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

func treeMinimum(n *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func treeMaximum(n *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func treeSuccessor(n *BinaryTreeNode) *BinaryTreeNode {
	if n.right != nil {
		return treeMinimum(n.right)
	} else {
		y := n.p
		for y != nil && n == y.right {
			n = y
			y = y.p
		}

		return y
	}
}

func treeInsert(T *BinaryTree, z *BinaryTreeNode) {
	x := T.root
	var y *BinaryTreeNode
	for x != nil {
		y = x
		if z.k < x.k {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nil {
		T.root = z
	} else if z.k < y.k {
		y.left = z
	} else {
		y.right = z
	}
}

func transplant(T *BinaryTree, u, v *BinaryTreeNode) {
	if u.p == nil {
		T.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	if v != nil {
		v.p = u.p
	}
}

func treeDelete(T *BinaryTree, z *BinaryTreeNode) {
	if z.left == nil {
		transplant(T, z, z.right)
	} else if z.right == nil {
		transplant(T, z, z.left)
	} else {
		y := treeMinimum(z.right)
		if y != z.right {
			transplant(T, y, y.right)
			y.right = z.right
			y.right.p = y
		}
		transplant(T, z, y)
		y.left = z.left
		y.left.p = y
	}
}

func main() {
	t := BinaryTree{}
	treeInsert(&t, &BinaryTreeNode{k: 1})
	treeInsert(&t, &BinaryTreeNode{k: 6})
	treeInsert(&t, &BinaryTreeNode{k: 3})
	treeInsert(&t, &BinaryTreeNode{k: 8})
	treeInsert(&t, &BinaryTreeNode{k: 5})
	treeInsert(&t, &BinaryTreeNode{k: 2})

	treeWalk(t.root)

	treeDelete(&t, &BinaryTreeNode{k: 1})
	treeDelete(&t, &BinaryTreeNode{k: 6})
	treeDelete(&t, &BinaryTreeNode{k: 3})
	treeDelete(&t, &BinaryTreeNode{k: 8})
	treeDelete(&t, &BinaryTreeNode{k: 5})
	treeDelete(&t, &BinaryTreeNode{k: 2})

	treeWalk(t.root)

}
