package main

import (
	"fmt"
	"morse-code/pkg/morse"
)

func main() {
	morseTree := morse.ConstructMorseTree()

	lookup_node := morseTree.FindNodeByCode(morseTree, "...-")
	fmt.Println("Lookup Value:", lookup_node.Value)

	lookup_node2 := morseTree.FindNodeByChar(morseTree, "A")
	fmt.Println("Lookup Code:", lookup_node2.Code)

	// if lookup_node.Dash != nil {
	// 	fmt.Println("Dash Value:", lookup_node.Dash.Value)
	// }
}
