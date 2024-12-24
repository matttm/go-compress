package bitreader

import "fmt"

type BitReader struct {
	encoded  []byte
	index    uint64 // index of current byte being examined in encoded
	current  uint8
	bitCount uint8 // the bit position being examined of byte pointed to by index
}

// Function FromSlice
// create a BitReader that will read oprivided slice
func FromSlice(s []byte) *BitReader {
	br := new(BitReader)
	br.encoded = s
	br.index = 0
	br.current = s[0]
	br.bitCount = 0 // todo: check if on big-endian or not
	return br
}

// Function ReadBit
// will read a single bit
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

// Function IsEOF
// Determines if there are more bits to read
func (br *BitReader) IsEOF() bool {
	return br.index >= uint64(len(br.encoded))
}

// Function IsLastByte
// Determines if the next bit to be read is in the last byte
func (br *BitReader) IsLastByte() bool {
	return br.index == uint64(len(br.encoded)-1)
}

// getBitPosition
// returns position of next bit to be read in current byte
func (br *BitReader) GetBitPosition() uint8 {
	return br.bitCount
}
