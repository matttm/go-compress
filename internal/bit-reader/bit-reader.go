package bitreader

import "fmt"

type BitReader struct {
	encoded  []byte
	index    uint64 // index of current byte being examined in encoded
	current  uint8
	bitCount uint8 // the bit position being examined of byte pointed to by index
}

func FromSlice(s []byte) *BitReader {
	br := new(BitReader)
	br.encoded = s
	br.index = 0
	br.current = s[0]
	br.bitCount = 0 // todo: check if on big-endian or not
	return br
}
func (br *BitReader) ReadBit() (bool, error) {
	if br.IsEOF() {
		return false, fmt.Errorf("Error: Index out-of-bounds")
	} else {
		if br.bitCount == 8 {
			br.bitCount = 0
			br.index += 1
			br.current = br.encoded[br.index]
			return br.ReadBit()
		} else {
			_set := br.current >> (7 - br.bitCount)
			fmt.Printf("%08b, %08b\n", _set, br.current)
			set := _set & 1
			br.bitCount += 1
			return set == 1, nil
		}
	}
}

func (br *BitReader) IsEOF() bool {
	return br.index == uint64(len(br.encoded))
}
