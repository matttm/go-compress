package bitwriter

import "fmt"

type BitWriter struct {
	storage  []byte
	current  uint8
	bitCount uint8
}

func WithSlice(b []byte) *BitWriter {
	bw := new(BitWriter)
	bw.storage = b
	bw.current = 0b0
	bw.bitCount = 0
	return bw
}

func (bw *BitWriter) WriteBit(set bool) {
	// the idea is well modify the lsb, then shift it left
	if set {
		bw.current |= 1
	}
	bw.bitCount += 1
	if bw.bitCount == 8 {
		bw.storage = append(bw.storage, byte(bw.current))
		bw.current = 0b0
		bw.bitCount = 0
	} else {
		bw.current <<= 1
	}
	fmt.Printf("%08b\n", bw.current)

}

func (bw *BitWriter) YieldSlice() []byte {
	if bw.bitCount > 0 {
		remaining := 8 - bw.bitCount
		for range remaining {
			bw.WriteBit(false)
		}
	}
	return bw.storage
}
