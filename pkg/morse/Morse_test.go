package morse

import (
	"testing"
)

// Test constructing the Morse tree
func TestConstructMorseTree(t *testing.T) {
	root := ConstructMorseTree()

	// Check if the root is not nil
	if root == nil {
		t.Fatal("Expected root to be non-nil")
	}

	// Check if the top-level Dot and Dash nodes exist
	if root.Dot == nil || root.Dash == nil {
		t.Fatalf("Expected root to have both Dot and Dash nodes, got Dot: %v, Dash: %v", root.Dot, root.Dash)
	}
}

// Test finding a node by its Morse code
func TestFindNodeByCode(t *testing.T) {
	root := ConstructMorseTree()

	// Test for a known Morse code (e.g., ".--" for 'W')
	node := root.FindNodeByCode(root, ".--")
	if node == nil {
		t.Fatal("Expected to find node for '.--', but got nil")
	}
	if node.Value != "W" {
		t.Fatalf("Expected node value 'W', but got '%s'", node.Value)
	}
}

// Test finding a node by its character (DFS)
func TestFindNodeByChar(t *testing.T) {
	root := ConstructMorseTree()

	// Test for character 'S' which has the Morse code "..."
	node := root.FindNodeByChar(root, "S")
	if node == nil {
		t.Fatal("Expected to find node for 'S', but got nil")
	}
	if node.Code != "..." {
		t.Fatalf("Expected Morse code '...' for 'S', but got '%s'", node.Code)
	}
}

// Test inserting a node into the tree
func TestInsertIntoTree(t *testing.T) {
	root := ConstructMorseTree()

	// Insert a new code (not in the original map) for testing purposes
	newCode := "---."
	value := "Test"
	inserted := insertIntoTree(root, newCode, value)

	// Ensure the insertion returns success
	if inserted != 1 {
		t.Fatalf("Expected 1 (successful insertion), but got %d", inserted)
	}

	// Verify the new node was inserted correctly
	insertedNode := root.FindNodeByCode(root, newCode)
	if insertedNode == nil {
		t.Fatal("Expected to find inserted node, but got nil")
	}

	// Verify the value of the inserted node
	if insertedNode.Value != value {
		t.Fatalf("Expected value '%s' for node, but got '%s'", value, insertedNode.Value)
	}
}
