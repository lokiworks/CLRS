package main

type Color int

const (
	RED Color = iota
	BLACK
)

type RedBlackTree struct {
	root *RedBlackNode
}

type RedBlackNode struct {
	right *RedBlackNode
	left  *RedBlackNode
	p     *RedBlackNode
	key   int
	color Color
}

func leftRotation(T *RedBlackTree, x *RedBlackNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.p = x
	}
	y.p = x.p
	if x.p == nil {
		T.root = y
	} else if x == x.p.left {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
}

func rightRotation(T *RedBlackTree, x *RedBlackNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.p = x
	}
	y.p = x.p
	if x.p == nil {
		T.root = y
	} else if x == x.p.right {
		x.p.right = y
	} else {
		x.p.left = y
	}
	y.right = x
	x.p = y
}

func rbInsert(T *RedBlackTree, z *RedBlackNode) {
	x := T.root
	var y *RedBlackNode
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nil {
		T.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	z.left, z.right = nil, nil
	z.color = RED
	rbInsertFixup(T, z)
}

func rbInsertFixup(T *RedBlackTree, z *RedBlackNode) {
	for z.p.color == RED {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					leftRotation(T, z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				rightRotation(T, z.p.p)
			}
		} else {
			y := z.p.p.left
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					rightRotation(T, z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				leftRotation(T, z.p.p)

			}
		}
	}

	T.root.color = BLACK
}

func rbTransplant(T *RedBlackTree, u, v *RedBlackNode) {
	if u.p == nil {
		T.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func treeMinimum(n *RedBlackNode) *RedBlackNode {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func rbDelete(T *RedBlackTree, z *RedBlackNode) {
	y := z
	yOriginalColor := y.color
	var x *RedBlackNode
	if z.left == nil {
		x = z.right
		rbTransplant(T, z, z.right) // replace z by its right child
	} else if z.right == nil {
		x = z.left
		rbTransplant(T, z, z.left)
	} else {
		y := treeMinimum(z.right) // y is z's successor
		yOriginalColor = y.color
		x = y.right
		if y != z.right {
			rbTransplant(T, y, y.right)
			y.right = z.right
			y.right.p = y
		} else {
			x.p = y
		}
		rbTransplant(T, z, y)
		y.left = z.left
		y.left.p = y
		y.color = z.color
	}
	// if any red-black violations occurred, correct them
	if yOriginalColor == BLACK {
		rbDeleteFixup(T, x)
	}

}

func rbDeleteFixup(T *RedBlackTree, x *RedBlackNode) {
	for x != T.root && x.color == BLACK {
		if x == x.p.left { // is x a left child
			w := x.p.right
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				leftRotation(T, x.p)
				w = x.p.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					rightRotation(T, w)
					w = x.p.right
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.right.color = BLACK
				leftRotation(T, x.p)
				x = T.root
			}
		} else {
			w := x.p.left
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				rightRotation(T, x.p)
				w = x.p.left
			}
			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					leftRotation(T, w)
					w = x.p.left
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.left.color = BLACK
				rightRotation(T, x.p)
				x = T.root
			}
		}
	}
	x.color = BLACK
}

func main() {
}
