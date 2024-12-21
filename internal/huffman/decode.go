package huffman

import (
	"bytes"
	"fmt"
	"strings"

	bitreader "github.com/matttm/go-compress/internal/bit-reader"
)

func FromEncodedText(s string) *HuffmanCodec {
	c := New()
	b := []byte(s)
	tree, index := deserializeTree(b)
	c.tree = tree
	c.encoded = b[index:]
	var sb strings.Builder
	c.createCodeTable(c.tree, sb)
	c.decode()
	return c
}

func (c *HuffmanCodec) decode() {
	var codeBldr strings.Builder
	br := bitreader.FromSlice(c.encoded)
	var decodedBldr strings.Builder
	// while br is not at eof
	for !br.IsEOF() {
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
	c.decoded = decodedBldr.String()
}
func deserializeTree(_data []byte) (*HuffmanNode, int) {
	if !bytes.HasPrefix(_data, MAGIC_NUMBER) {
	}
	// count := data[0]
	//
	// data[0:2] is magic
	index := 0
	data := _data[2:]
	var dfs func() *HuffmanNode
	dfs = func() *HuffmanNode {
		if index >= len(data) {
			return nil
		}
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
	// printTree(n)
	return n, index + 2 // add 2 to skip magic
}
