package main

// FGKNode represents a node in the FGK tree.
type FGKNode struct {
	character  rune
	frequency  int
	left       *FGKNode
	right      *FGKNode
	parent     *FGKNode
	isInternal bool
}

// FGKTree represents the adaptive Huffman tree.
type FGKTree struct {
	root      *FGKNode
	leafNodes map[rune]*FGKNode
	zeroNode  *FGKNode
}

// NewFGKTree initializes a new FGK tree.
func NewFGKTree() *FGKTree {
	zeroNode := &FGKNode{frequency: 0, isInternal: true}
	return &FGKTree{
		root:      zeroNode,
		leafNodes: make(map[rune]*FGKNode),
		zeroNode:  zeroNode,
	}
}

// UpdateTree updates the FGK tree with the given character.
func (tree *FGKTree) UpdateTree(character rune) {
	if node, exists := tree.leafNodes[character]; exists {
		// Increment frequency of existing character
		tree.incrementFrequency(node)
		return
	}

	// Create a new node for the character
	newNode := &FGKNode{
		character:  character,
		frequency:  1,
		isInternal: false,
	}
	
	// Insert the new node
	tree.insertNode(newNode)
}

// incrementFrequency increments the frequency of the given node and updates the tree.
func (tree *FGKTree) incrementFrequency(node *FGKNode) {
	node.frequency++
	
	// Update the frequency of all ancestor nodes
	for node.parent != nil {
		node = node.parent
		node.frequency++
	}

	// Now update the tree to maintain FGK properties
	tree.updateTree()
}

// insertNode inserts a new node into the tree.
func (tree *FGKTree) insertNode(node *FGKNode) {
	// Insert node and adjust tree
	// Implementation details would include finding the appropriate position
	// and restructuring the tree to maintain the FGK properties.
	tree.leafNodes[node.character] = node
	tree.updateTree()
}

// updateTree updates the FGK tree by restructuring it if necessary.
func (tree *FGKTree) updateTree() {
	// Restructure tree to maintain FGK properties
	// This will involve rotations and updates to maintain the FGK structure.
}

// Encode encodes a string using the FGK tree.
func (tree *FGKTree) Encode(input string) string {
	encoded := ""
	for _, char := range input {
		tree.UpdateTree(char)
		encoded += tree.getCodeForCharacter(char)
	}
	return encoded
}

// getCodeForCharacter retrieves the Huffman code for a given character.
func (tree *FGKTree) getCodeForCharacter(character rune) string {
	node, exists := tree.leafNodes[character]
	if !exists {
		return ""
	}
	
	var code string
	for node.parent != nil {
		if node == node.parent.left {
			code = "0" + code
		} else {
			code = "1" + code
		}
		node = node.parent
	}
	return code
}

// Decode decodes a string using the FGK tree.
func (tree *FGKTree) Decode(encoded string) string {
	result := ""
	node := tree.root
	for _, bit := range encoded {
		if bit == '0' {
			node = node.left
		} else {
			node = node.right
		}
		if node.left == nil && node.right == nil {
			result += string(node.character)
			node = tree.root
		}
	}
	return result
}
