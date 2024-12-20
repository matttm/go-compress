package huffman

import (
	"bytes"
	"strings"

	bitreader "github.com/matttm/go-compress/internal/bit-reader"
)

func (c *HuffmanCodec) FromEncodedText(s string) {
}

func (c *HuffmanCodec) decode() {
	var codeBldr strings.Builder
	br := bitreader.FromSlice(c.encoded)
	var decodedBldr strings.Builder
	// while br is not at eof
	for !br.IsEOF() {
		_bit, _ := br.ReadBit()
		var b rune
		if _bit {
			b = '1'
		} else {
			b = '0'
		}
		codeBldr.WriteRune(b)
		code := codeBldr.String()
		if r, ok := c.decodingTable[code]; ok {
			continue
		} else {
			decodedBldr.WriteRune(r)
		}
	}
	c.decoded = decodedBldr.String()
}
func deserializeTree(_data []byte) *HuffmanNode {
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
	return n
}
