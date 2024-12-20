package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type huffman_testcase struct {
	input    string
	expected []byte
}
type serialize_testcase struct {
	createTree func() *HuffmanNode
	expected   []byte
}

func Test_ShouldEncode(t *testing.T) {
	tests := []huffman_testcase{
		{
			input:    "aabba",
			expected: []byte{0b11001000},
		},
		{
			input:    "aabbaaabb",
			expected: []byte{0b11001110, 0b00000000},
		},
		{
			input:    "aabbaaaba",
			expected: []byte{0b11001110, 0b10000000},
		},
	}
	for _, _t := range tests {
		c := FromDecodedText(_t.input)
		assert.ElementsMatch(t, c.encoded, _t.expected, "decode")
	}
}
func Test_ShouldSerialize(t *testing.T) {
	tests := []serialize_testcase{
		{
			createTree: func() *HuffmanNode {
				a := HuffmanNode{symbol: '*'}
				b := HuffmanNode{symbol: 'b'}
				c := HuffmanNode{symbol: 'c'}
				a.left = &b
				a.right = &c
				return &a
			},
			expected: []byte{'*', 'b', NULL, NULL, 'c', NULL, NULL},
		},
		{
			createTree: func() *HuffmanNode {
				a1 := HuffmanNode{symbol: '*'}
				a2 := HuffmanNode{symbol: '*'}
				b := HuffmanNode{symbol: 'b'}
				c := HuffmanNode{symbol: 'c'}
				a1.left = &b
				a1.right = &a2
				a2.left = &c
				a2.right = nil
				return &a1
			},
			expected: []byte{'*', 'b', NULL, NULL, '*', 'c', NULL, NULL, NULL},
		},
	}
	for _, _t := range tests {
		_expected := MAGIC_NUMBER
		_expected = append(_expected, _t.expected...)
		assert.Equal(t, serializeTree(_t.createTree()), _expected, "serialize basic tree")
	}
}
