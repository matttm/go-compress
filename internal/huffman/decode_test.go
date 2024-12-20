package huffman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type decode_testcase struct {
	input    []byte
	expected string
}
type deserialize_testcase struct {
	expectedTree *HuffmanNode
	input        []byte
}

//	func Test_ShouldDecode(t *testing.T) {
//		tests := []decode_testcase{
//			{
//				expected: "aabba",
//				input:    []byte{0b11001000},
//			},
//		}
//		for _, _t := range tests {
//			c := FromEncodedText(_t.input)
//			assert.ElementsMatch(t, c.encoded, _t.expected, "decode")
//		}
//	}
func Test_ShouldDeserialize(t *testing.T) {
	var compareTrees func(t *testing.T, a, b *HuffmanNode) bool
	compareTrees = func(t *testing.T, a, b *HuffmanNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		return a.symbol == b.symbol && compareTrees(t, a.left, b.left) && compareTrees(t, a.right, b.right)
	}
	tests := []deserialize_testcase{
		{
			expectedTree: func() *HuffmanNode {
				a := HuffmanNode{symbol: '*'}
				b := HuffmanNode{symbol: 'b'}
				c := HuffmanNode{symbol: 'c'}
				a.left = &b
				a.right = &c
				return &a
			}(),
			input: []byte{'*', 'b', NULL, NULL, 'c', NULL, NULL},
		},
	}
	for _, _t := range tests {
		_input := MAGIC_NUMBER
		_input = append(_input, _t.input...)
		treeA := _t.expectedTree
		treeB := deserializeTree(_input)
		fmt.Println("tree  -------------")
		printTree(treeA)
		fmt.Println("tree  -------------")
		printTree(treeB)
		assert.Equal(t, compareTrees(t, treeA, treeB), true, "comparing trees")
	}
}
