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
