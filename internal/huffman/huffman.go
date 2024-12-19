package huffman

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
