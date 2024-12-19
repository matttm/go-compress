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
func deserializeTree(data []byte) *HuffmanNode {
	if !bytes.HasPrefix(data, MAGIC_NUMBER) {
	}
	n := new(HuffmanNode)
	for _, d := range data {
		switch d {
		// left node
		case 0:
			break
		case 1:
			break
		default:
			return
		}
	}
	return n
}
