package morse

import (
	"fmt"
	"sort"
)

type MorseNode struct {
	Code  string
	Value string
	Dot   *MorseNode
	Dash  *MorseNode
}

var MorseToEnglish = map[string]string{
	".":    "E",
	"-":    "T",
	"..":   "I",
	".-":   "A",
	"-.":   "N",
	"--":   "M",
	"...":  "S",
	"..-":  "U",
	".-.":  "R",
	".--":  "W", // Mapping for '.--' to 'W'
	"-..":  "D",
	"-.-":  "K",
	"--.":  "G",
	"---":  "O",
	"....": "H",
	"...-": "V",
	"..-.": "F",
	".-..": "L",
	".--.": "P",
	".---": "J",
	"-...": "B",
	"-..-": "X",
	"-.-.": "C",
	"-.--": "Y",
	"--..": "Z",
	"--.-": "Q",
}

/* Methods */

// Method to set the value of a MorseNode
func (node *MorseNode) SetValue(value string) {
	node.Value = value
}

// Method to add a dot node
func (node *MorseNode) AddDot(dotNode *MorseNode) {
	node.Dot = dotNode
}

// Method to add a dash node
func (node *MorseNode) AddDash(dashNode *MorseNode) {
	node.Dash = dashNode
}

// Get next dot (if exists)
func (node *MorseNode) GetNextDot() *MorseNode {
	return node.Dot
}

func (node *MorseNode) GetNextDash() *MorseNode {
	return node.Dash
}

func (node *MorseNode) FindNodeByCode(cursorNode *MorseNode, code string) *MorseNode {
	// cursorDepth
	cursorDepth := len(cursorNode.Code)

	//found it
	if cursorNode.Code == code {
		return cursorNode
	}

	isDot := code[cursorDepth] == '.'

	if isDot {
		return cursorNode.Dot.FindNodeByCode(cursorNode.Dot, code)
	}
	return cursorNode.Dash.FindNodeByCode(cursorNode.Dash, code)

}

// perform a DFS
func (node *MorseNode) FindNodeByChar(cursorNode *MorseNode, val string) *MorseNode {
	// Base case: Check if the current node's value matches the target value
	if cursorNode == nil {
		return nil // Reached a leaf node, so return nil
	}

	if cursorNode.Value == val {
		// Found the node with the matching value
		return cursorNode
	}

	// Perform DFS on the Dot (left) node
	if foundNode := cursorNode.FindNodeByChar(cursorNode.Dot, val); foundNode != nil {
		return foundNode
	}

	// If not found in Dot, perform DFS on the Dash (right) node
	return cursorNode.FindNodeByChar(cursorNode.Dash, val)
}

// FindNodeByCharBFS performs a breadth-first search (BFS) to find a node by its value
func (node *MorseNode) FindNodeByCharBFS(cursorNode *MorseNode, val string) *MorseNode {
	// Base case: if the tree is empty
	if cursorNode == nil {
		return nil
	}

	// Initialize a queue with the root node
	queue := []*MorseNode{cursorNode}

	// While the queue is not empty, keep searching
	for len(queue) > 0 {
		// Dequeue the first node in the queue
		currentNode := queue[0]
		queue = queue[1:]

		fmt.Printf("FindNodeByCharBFS\tcode:%s\tvalue:%s\n", currentNode.Code, currentNode.Value)

		// Check if the current node's value matches the target value
		if currentNode.Value == val {
			return currentNode
		}

		// Add the Dot (left) node to the queue if it exists
		if currentNode.Dot != nil {
			queue = append(queue, currentNode.Dot)
		}

		// Add the Dash (right) node to the queue if it exists
		if currentNode.Dash != nil {
			queue = append(queue, currentNode.Dash)
		}
	}

	// If we exhaust the queue and haven't found the node, return nil
	return nil
}

//utils

func ConstructMorseTree() *MorseNode {
	var rootNode MorseNode = MorseNode{Code: "", Value: ""}
	//	var cursor MorseNode
	var morseCodes []string
	for morseCode := range MorseToEnglish {
		morseCodes = append(morseCodes, morseCode)
	}

	// Step 2: Sort the keys by length of the Morse code
	sort.Slice(morseCodes, func(i, j int) bool {
		return len(morseCodes[i]) < len(morseCodes[j])
	})

	// Step 3: Iterate over the sorted Morse codes and access their corresponding English character
	for _, morseCode := range morseCodes {
		insertIntoTree(&rootNode, morseCode, MorseToEnglish[morseCode])
	}

	return &rootNode
}

func insertIntoTree(cursorNode *MorseNode, morseCode string, value string) int {
	//starts at 0
	currentDepth := len(cursorNode.Code)
	morseCodeLength := len(morseCode)

	isDot := morseCode[currentDepth] == '.'

	//if code is larger than the currentdepth, insert said nodes
	if morseCodeLength-1 == currentDepth {
		if isDot {
			cursorNode.Dot = &MorseNode{Code: morseCode, Value: value}
			return 1
		} else {
			cursorNode.Dash = &MorseNode{Code: morseCode, Value: value}
			return 1
		}
	}

	//traverse
	if morseCodeLength > currentDepth {
		if isDot {
			return insertIntoTree(cursorNode.Dot, morseCode, value)
		}
		if !isDot {
			return insertIntoTree(cursorNode.Dash, morseCode, value)
		}
	}

	return 1
}
