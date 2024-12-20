package huffman

import (
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
	var compareTrees func(t *testing.T, a, b *HuffmanNode)
	compareTrees = func(t *testing.T, a, b *HuffmanNode) {
		assert.Equal(t, serializeTree(_t.createTree()), _input, "deserialize basic tree")
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
		compareTrees(t, _t.expectedTree, deserializeTree(_t.input))
	}
}
