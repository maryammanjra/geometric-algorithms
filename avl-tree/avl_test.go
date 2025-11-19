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
