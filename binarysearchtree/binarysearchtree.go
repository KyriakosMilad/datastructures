package binarysearchtree

import (
	"fmt"
	"github.com/KyriakosMilad/datastructures/queue"
)

type node struct {
	value int
	left  *node
	right *node
	count int
}

func (n *node) delete() {
	n.left = nil
	n.right = nil
	n.value = 0
	n.count = 0
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

/** todo
* clean code !!
* doing it recursively would be a lot better
 */

func (bst *BinarySearchTree) Remove(val int) error {
	err := fmt.Errorf("no node found with value %v", val)
	if bst.root == nil {
		return err
	}

	parent := bst.root
	current := bst.root
	for current != nil {
		if val < current.value {
			parent = current
			current = current.left
		} else if val > current.value {
			parent = current
			current = current.right
		} else {
			if current.left != nil && current.right != nil {
				minParent := current
				min := current.right
				for min.left != nil {
					minParent = min
					min = min.left
				}
				if min.right != nil {
					minParent.left = min.right
				} else {
					if minParent.left == min {
						minParent.left = nil
					} else if minParent.right == min {
						minParent.right = nil
					}
				}
				current.value = min.value
				current.count = min.count
				min.delete()
			} else if current.left == nil && current.right != nil {
				if parent.right == current {
					parent.right = current.right
					current.delete()
				} else if parent.left == current {
					parent.left = current.right
					current.delete()
				} else {
					bst.root = current.right
					current.delete()
				}
			} else if current.right == nil && current.left != nil {
				if parent.right == current {
					parent.right = current.left
					current.delete()
				} else if parent.left == current {
					parent.left = current.left
					current.delete()
				} else {
					bst.root = current.left
					current.delete()
				}
			} else if current.left == nil && current.right == nil {
				if current == bst.root {
					bst.root = nil
				} else if parent.right == current {
					parent.right = nil
					current.delete()

				} else if parent.left == current {
					parent.left = nil
					current.delete()
				}
			}
			return nil
		}
	}

	return err
}

func (bst *BinarySearchTree) Contains(val int) bool {
	if bst.root == nil {
		return false
	}

	current := bst.root
	for current != nil {
		if val == current.value {
			return true
		} else if val > current.value {
			current = current.right
		} else if val < current.value {
			current = current.left
		}
	}

	return false
}

func (bst *BinarySearchTree) BreadthFirstSearch() []int {
	if bst.root == nil {
		return []int{}
	}

	q := queue.New(bst.root)
	visited := []int{}

	for q.Size() > 0 {
		_, v := q.First()
		if n, ok := v.(*node); ok {
			visited = append(visited, n.value)

			if n.right != nil {
				q.Enqueue(n.right)
			}
			if n.left != nil {
				q.Enqueue(n.left)
			}

			q.Dequeue()
		} else {
			panic(fmt.Errorf("%v is not a node", v))
		}
	}

	return visited
}

var visited []int

func preOrderVisit(n *node) {
	visited = append(visited, n.value)
	if n.right != nil {
		preOrderVisit(n.right)
	}
	if n.left != nil {
		preOrderVisit(n.left)
	}
}

func (bst *BinarySearchTree) PreOrderTraversal() []int {
	if bst.root == nil {
		return []int{}
	}
	preOrderVisit(bst.root)
	return visited
}
