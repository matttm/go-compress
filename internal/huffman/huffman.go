package huffman

import "strings"

var MAGIC_NUMBER = []byte{0x80, 0xF0}
var NULL byte = '@'

type HuffmanCodec struct {
	frequencyTable map[rune]int
	encodingTable  map[rune]string // maps rune to code
	decodingTable  map[string]rune // maps code to rune
	tree           *HuffmanNode
	encoded        []byte
	decoded        string
}

func New() *HuffmanCodec {
	c := new(HuffmanCodec)
	c.encodingTable = make(map[rune]string)
	c.decodingTable = make(map[string]rune)
	return c
}

func (c *HuffmanCodec) createCodeTable(node *HuffmanNode, sb strings.Builder) {
	if node == nil {
		return
	}
	if node.symbol != '*' {
		s := sb.String()
		symbol := node.symbol
		c.encodingTable[symbol] = s
		c.decodingTable[s] = symbol
		return
	}
	if node.left != nil {
		var left strings.Builder
		left.WriteString(sb.String())
		left.WriteRune('0')
		c.createCodeTable(node.left, left)
	}
	if node.right != nil {
		var right strings.Builder
		right.WriteString(sb.String())
		right.WriteRune('1')
		c.createCodeTable(node.right, right)
	}
}
