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

func Test_ShouldDecode(t *testing.T) {
	tests := []decode_testcase{
		{
			expected: "aabba",
			// TODO: devise a way to indicate how many bits to ignore (i.e. gotta to ignore the laast 3 bits in byte below)
			input: []byte{'*', 'a', NULL, NULL, 'b', NULL, NULL, 0b11001000},
		},
	}
	for _, _t := range tests {
		_input := MAGIC_NUMBER
		_input = append(_input, _t.input...)
		c := FromEncodedText(string(_input))
		assert.Equal(t, c.decoded, _t.expected, "decode")
	}
}
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
		treeB, _ := deserializeTree(_input)
		assert.Equal(t, true, compareTrees(t, treeA, treeB), "comparing trees")
	}
}
