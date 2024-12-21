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
		br.current = br.encoded[br.index]
		if br.bitCount == 8 {
			br.bitCount = 0
			br.index += 1
			return br.ReadBit()
		} else {
			_set := br.current >> (7 - br.bitCount)
			set := _set & 1
			br.bitCount += 1
			return set == 1, nil
		}
	}
}

func (br *BitReader) IsEOF() bool {
	return br.index >= uint64(len(br.encoded))
}
func (br *BitReader) IsLastByte() bool {
	return br.index == uint64(len(br.encoded)-1)
}

// getBitPosition
// returns position of next bit to be read in current byte
func (br *BitReader) GetBitPosition() uint8 {
	return br.bitCount
}
