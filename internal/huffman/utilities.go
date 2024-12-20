package huffman

import "fmt"

func CreateTestTreeFromSlice(_symbols string) *HuffmanNode {
	//  symbols := []rune(_symbols)
	return nil
}
func printTree(node *HuffmanNode) {
	if node == nil {
		return
	}
	fmt.Printf("%c\n", node.symbol)
	printTree(node.left)
	printTree(node.right)

}
