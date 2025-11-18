package avl

type Comparator[T any] func(a, b T) int

type Node[T any] struct {
	Value   T
	Height  int
	Balance int
	Left    *Node[T]
	Right   *Node[T]
}

type AVLTree[T any] struct {
	Root       *Node[T]
	Comparator Comparator[T]
}

func NewAVL[T any](cmp func(a, b T) int) *AVLTree[T] {
	avl := AVLTree[T]{Comparator: cmp}
	return &avl
}

func SearchTree[T any](tree *AVLTree[T], val T) bool {
	if tree == nil {
		return false
	}

	currNode := tree.Root
	for currNode != nil {

		difference := tree.Comparator(currNode.Value, val)
		if difference == 0 {
			return true
		}

		if difference > 0 {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
	}

	return false
}

func Insert[T any](tree *AVLTree[T], val T) {
	newNode := &Node[T]{val, 0, 0, nil, nil}

	if tree == nil {
		tree.Root = newNode
		return
	}

	currNode := tree.Root
	parent := currNode

	for currNode != nil {
		parent = currNode
		difference := tree.Comparator(currNode.Value, val)

		if difference > 0 {
			currNode = currNode.Right
		} else {
			currNode = currNode.Left
		}
	}

	newNode.Height = parent.Height + 1

	if tree.Comparator(parent.Value, val) > 0 {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	parent.Balance = parent.Left.Height - parent.Right.Height
	return
}
