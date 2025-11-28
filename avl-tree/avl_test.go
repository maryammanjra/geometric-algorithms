package avl

import "testing"

func intCmp(a, b int) int { return a - b }

func TestInsertAndSearch(t *testing.T) {
	tree := NewAVL(intCmp)

	values := []int{10, 20, 5, 4, 15, 30}
	for _, v := range values {
		tree.Root = Insert(tree.Root, tree.Comparator, v)
	}

	tests := []struct {
		val      int
		expected bool
	}{
		{10, true},
		{20, true},
		{4, true},
		{15, true},
		{99, false},
	}

	for _, tc := range tests {
		result := SearchTree(tree, tc.val)
		if result != tc.expected {
			t.Errorf("Search(%d) = %v, want %v", tc.val, result, tc.expected)
		}
	}
}

func TestBalancing(t *testing.T) {
	tree := NewAVL(intCmp)

	for _, v := range []int{1, 2, 3, 4, 5} {
		tree.Root = Insert(tree.Root, tree.Comparator, v)
	}

	if tree.Root.Value != 2 && tree.Root.Value != 3 {
		t.Errorf("Tree is not balanced. Root = %v", tree.Root.Value)
	}

	if tree.Root.Balance < -1 || tree.Root.Balance > 1 {
		t.Errorf("Invalid root balance factor: %d", tree.Root.Balance)
	}
}

// --- New deletion tests ---

func TestDeleteLeafNode(t *testing.T) {
	tree := NewAVL(intCmp)
	values := []int{10, 20, 5}
	for _, v := range values {
		tree.Root = Insert(tree.Root, tree.Comparator, v)
	}

	tree.Root = Delete(tree.Root, tree.Comparator, 5)

	if SearchTree(tree, 5) {
		t.Errorf("Deleted value 5 still found in tree")
	}

	checkAVLProperties(t, tree.Root)
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	tree := NewAVL(intCmp)
	values := []int{10, 5, 20, 15}
	for _, v := range values {
		tree.Root = Insert(tree.Root, tree.Comparator, v)
	}

	tree.Root = Delete(tree.Root, tree.Comparator, 20)

	if SearchTree(tree, 20) {
		t.Errorf("Deleted value 20 still found in tree")
	}

	checkAVLProperties(t, tree.Root)
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	tree := NewAVL(intCmp)
	values := []int{10, 5, 20, 15, 25}
	for _, v := range values {
		tree.Root = Insert(tree.Root, tree.Comparator, v)
	}

	tree.Root = Delete(tree.Root, tree.Comparator, 20)

	if SearchTree(tree, 20) {
		t.Errorf("Deleted value 20 still found in tree")
	}

	checkAVLProperties(t, tree.Root)
}

func TestDeleteRootNode(t *testing.T) {
	tree := NewAVL(intCmp)
	values := []int{10, 5, 15}
	for _, v := range values {
		tree.Root = Insert(tree.Root, tree.Comparator, v)
	}

	tree.Root = Delete(tree.Root, tree.Comparator, 10)

	if SearchTree(tree, 10) {
		t.Errorf("Deleted root value 10 still found in tree")
	}

	checkAVLProperties(t, tree.Root)
}

// --- Helper to check AVL balance recursively ---
func checkAVLProperties(t *testing.T, node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := checkAVLProperties(t, node.Left)
	rightHeight := checkAVLProperties(t, node.Right)

	balance := leftHeight - rightHeight
	if balance < -1 || balance > 1 {
		t.Errorf("Node %v has invalid balance %d", node.Value, balance)
	}

	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
