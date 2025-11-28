package avl

import "fmt"

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

	//Inserted into left child of right subtree
	if (node.Balance < -1) && (getBalance(node.Right) > 0) {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	if (node.Balance > 1) && (getBalance(node.Left) < 0) {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	return node
}

func Delete[T any](node *Node[T], cmp Comparator[T], val T) *Node[T] {

	if node == nil {
		return nil
	}

	difference := cmp(node.Value, val)

	if difference == 0 {

		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

	}

	if difference > 0 {
		node.Left = Delete(node.Left, cmp, val)
	} else if difference < 0 {
		node.Right = Delete(node.Right, cmp, val)
	}

	node.Height = 1 + max(getHeight(node.Right), getHeight(node.Left))
	node.Balance = getBalance(node)

	if (node.Balance > 1) && (getBalance(node.Left) >= 0) {
		return rotateRight(node)
	}

	if (node.Balance < -1) && (getBalance(node.Right) <= 0) {
		return rotateLeft(node)
	}

	if (node.Balance < -1) && (getBalance(node.Right) > 0) {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	if (node.Balance > 1) && (getBalance(node.Left) < 0) {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	return node
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

func (t *AVLTree[T]) PrintGraphical() {
	if t.Root == nil {
		println("(empty)")
		return
	}
	printGraphicalNode(t.Root, "", true)
}

func printGraphicalNode[T any](node *Node[T], prefix string, isTail bool) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		printGraphicalNode(node.Right, newPrefix, false)
	}

	connector := "└── "
	if !isTail {
		connector = "┌── "
	}

	println(prefix + connector + fmt.Sprintf("%v (h=%d,b=%d)", node.Value, node.Height, node.Balance))

	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		printGraphicalNode(node.Left, newPrefix, true)
	}
}
