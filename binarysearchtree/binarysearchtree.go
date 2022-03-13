package binarysearchtree

type node struct {
	value int
	left  *node
	right *node
	count int
}

type BinarySearchTree struct {
	root *node
}

func New(n *node) *BinarySearchTree {
	return &BinarySearchTree{root: n}
}

func (bst *BinarySearchTree) Insert(val int) {
	if bst.root == nil {
		bst.root = &node{value: val}
	} else {
		current := bst.root
		for true {
			if val == current.value {
				current.count++
				break
			} else if val < current.value {
				if current.left == nil {
					current.left = &node{value: val}
					break
				} else {
					current = current.left
				}
			} else if val > current.value {
				if current.right == nil {
					current.right = &node{value: val}
					break
				} else {
					current = current.right
				}
			}
		}
	}
}
