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

func main() {
}
