package avl

import (
	"fmt"
	"testing"
)

func IntComparator(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestInsertAndSearch(t *testing.T) {
	tree := NewAVL(IntComparator)

	tree.Root = Insert(tree.Root, tree.Comparator, 10)
	tree.Root = Insert(tree.Root, tree.Comparator, 20)
	tree.Root = Insert(tree.Root, tree.Comparator, 30)

	tests := []struct {
		val  int
		want bool
	}{
		{10, true},
		{20, true},
		{30, true},
		{40, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Search %d", tt.val), func(t *testing.T) {
			got := SearchTree(tree, tt.val)
			if got != tt.want {
				t.Errorf("SearchTree(%d) = %v; want %v", tt.val, got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tree := NewAVL(IntComparator)

	tree.Root = Insert(tree.Root, tree.Comparator, 10)
	tree.Root = Insert(tree.Root, tree.Comparator, 20)
	tree.Root = Insert(tree.Root, tree.Comparator, 30)

	tree.Root = Delete(tree.Root, tree.Comparator, 20)

	if SearchTree(tree, 20) {
		t.Errorf("Expected 20 to be deleted")
	}

	if !SearchTree(tree, 10) {
		t.Errorf("Expected 10 to be found")
	}
	if !SearchTree(tree, 30) {
		t.Errorf("Expected 30 to be found")
	}
}

func TestAVLTreeBalance(t *testing.T) {
	tree := NewAVL(IntComparator)

	insertValues := []int{10, 20, 30, 25, 5, 15}
	for _, val := range insertValues {
		tree.Root = Insert(tree.Root, tree.Comparator, val)
	}

	tests := []struct {
		val      int
		expectBF int
	}{
		{5, 0},  // 5 has no children, so balance factor is 0
		{10, 0}, // 10 has left child 5 and right child 15, balance factor is 0
		{15, 0}, // 15 has no children, so balance factor is 0
		{20, 0}, // 20 has left child 10 and right child 30, balance factor is 0
		{25, 0}, // 25 has no children, so balance factor is 0
		{30, 1}, // 30 has left child 25 and no right child, balance factor is 1
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Check Balance Factor for %d", tt.val), func(t *testing.T) {
			node := searchNode(tree.Root, tree.Comparator, tt.val)
			if node != nil {
				balanceFactor := getBalance(node)
				if balanceFactor != tt.expectBF {
					t.Errorf("Balance factor for %d is %d, expected %d", tt.val, balanceFactor, tt.expectBF)
				}
			} else {
				t.Errorf("Node %d not found in the tree", tt.val)
			}
		})
	}
}

func searchNode[T any](node *Node[T], cmp Comparator[T], val T) *Node[T] {
	if node == nil {
		return nil
	}

	comparison := cmp(node.Value, val)
	if comparison == 0 {
		return node
	}

	if comparison > 0 {
		return searchNode(node.Left, cmp, val)
	} else {
		return searchNode(node.Right, cmp, val)
	}
}
