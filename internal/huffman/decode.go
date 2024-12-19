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
	// data[0:2] is magic
	// data[2:3] is count
	_data = _data[3:]
	var dfs func([]byte) *HuffmanNode
	dfs = func(data []byte) *HuffmanNode {
		n := new(HuffmanNode)
		n.symbol = '*'
		d := data[0]
		switch d {
		// left node
		case 0:
			n.left = dfs(data[1:])
			break
		// right node
		case 1:
			n.right = dfs(data[1:])
			break
		case NULL:
			return nil
		// must be a symbolic-leaf
		default:
			n.symbol = rune(d)
		}
		return n
	}
	return dfs(_data)
}
