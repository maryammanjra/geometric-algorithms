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

func getHeight[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}

	return node.Height
}

func getBalance[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}

	return getHeight(node.Left) - getHeight(node.Right)
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

func Insert[T any](node *Node[T], cmp Comparator[T], val T) *Node[T] {

	if node == nil {
		return &Node[T]{val, 1, 0, nil, nil}
	}

	difference := cmp(node.Value, val)

	if difference > 0 {
		node.Left = Insert(node.Left, cmp, val)
	} else if difference < 0 {
		node.Right = Insert(node.Right, cmp, val)
	}

	node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))
	node.Balance = getBalance(node)

	// Inserted into left child of left-subtree where left child may have had right child to begin with, or not
	if (node.Balance > 1) && (getBalance(node.Left) >= 0) {
		return rotateRight(node)
	}

	//Inserted into right child of right-subtree where right child may have had left child thus balance may equal 0
	if (node.Balance < -1) && (getBalance(node.Right) <= 0) {
		return rotateLeft(node)
	}

	return nil
}

func rotateRight[T any](node *Node[T]) *Node[T] {
	tmp := node.Left
	node.Left = tmp.Right
	tmp.Right = node
	return tmp
}

func rotateLeft[T any](node *Node[T]) *Node[T] {
	tmp := node.Right
	node.Right = tmp.Left
	tmp.Left = node
	return tmp
}
