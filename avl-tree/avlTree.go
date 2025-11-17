package avl

type Comparator[T any] func(a, b T) int

type Node[T any] struct {
	Value  T
	Height int
	Left   *Node[T]
	Right  *Node[T]
}

type AVLTree[T any] struct {
	root       *Node[T]
	comparator Comparator[T]
}
