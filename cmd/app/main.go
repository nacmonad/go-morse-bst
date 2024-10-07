package main

import (
	"fmt"
	"morse-code/pkg/morse"
)

func main() {
	morseTree := morse.ConstructMorseTree()

	lookup_code := "...-"
	lookup_node := morseTree.FindNodeByCode(morseTree, lookup_code)
	fmt.Printf("Lookup Value for %s: \t %s\n", lookup_code, lookup_node.Value)

	lookup_char := "E"
	lookup_node2 := morseTree.FindNodeByChar(morseTree, lookup_char)
	fmt.Printf("Lookup Code for %s: \t%s\n", lookup_char, lookup_node2.Code)

	lookup_char3 := "Z"
	lookup_node3 := morseTree.FindNodeByCharBFS(morseTree, lookup_char3)
	fmt.Printf("Lookup (BFS) Code for %s: \t%s\n", lookup_char3, lookup_node3.Code)

	// if lookup_node.Dash != nil {
	// 	fmt.Println("Dash Value:", lookup_node.Dash.Value)
	// }
}
