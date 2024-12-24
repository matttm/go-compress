package huffman

import (
	"fmt"
	"strings"

	bitreader "github.com/matttm/go-compress/internal/bit-reader"
)

// Function FromEncodedText
//
//	reads encoded text and output the decoded text to stdin
func FromEncodedText(s string) *HuffmanCodec {
	c := New()
	b := []byte(s)
	_, extra, data := unpackageData(b)
	tree, index := deserializeTree(data)
	c.tree = tree
	c.encoded = data[index:]
	var sb strings.Builder
	c.createCodeTable(c.tree, sb)
	c.decoded = c.decode(extra)
	fmt.Println(c.decoded)
	return c
}

// Function decode
//
//	decodes the internal field encoded using the field tree of *HuffmanNode
func (c *HuffmanCodec) decode(extra uint8) string {
	var codeBldr strings.Builder
	br := bitreader.FromSlice(c.encoded)
	var decodedBldr strings.Builder
	// while br is not at eof
	for !br.IsEOF() {
		if br.IsLastByte() && 8-br.GetBitPosition() == extra {
			break
		}
		_bit, err := br.ReadBit()
		if err != nil {
			break
		}
		var b rune
		if _bit {
			b = '1'
		} else {
			b = '0'
		}
		codeBldr.WriteRune(b)
		code := codeBldr.String()
		if r, ok := c.decodingTable[code]; !ok {
			fmt.Println(code)
			continue
		} else {
			decodedBldr.WriteRune(r)
			codeBldr.Reset()
		}
	}
	return decodedBldr.String()
}

// Function deserializeTree
//
//	takes binary data and creates a huffman tree from the beginning of it
//
//	the returned int is the index following the final serialized node, MEANING,
//	it dictates where the ACTUAL data to be decoded starts
func deserializeTree(data []byte) (*HuffmanNode, int) {
	// count := data[0]
	//
	index := 0
	var dfs func() *HuffmanNode
	dfs = func() *HuffmanNode {
		n := new(HuffmanNode)
		d := data[index]
		index += 1
		if d == NULL {
			return nil
		}
		n.symbol = rune(d)
		n.left = dfs()
		n.right = dfs()
		return n
	}
	n := dfs()
	return n, index
}
func unpackageData(b []byte) ([]byte, uint8, []byte) {
	header := b[:2]
	extra := uint8(b[2])
	return header, extra, b[3:]
}
