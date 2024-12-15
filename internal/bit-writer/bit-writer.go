package bitwriter

import "fmt"

type BitWriter struct {
	storage  []byte
	current  uint8
	bitIndex uint8
}

func WithSlice(b []byte) *BitWriter {
	bw := new(BitWriter)
	bw.storage = b
	bw.current = 0b0
	bw.bitIndex = 0
	return bw
}

func (bw *BitWriter) WriteBit(set bool) {
	// the idea is well modify the lsb, then shift it left
	fmt.Printf("%08b\n", bw.current)
	if set {
		bw.current |= 1
	}
	bw.current <<= 1
	bw.bitIndex += 1
	if bw.bitIndex == 7 {
		bw.storage = append(bw.storage, byte(bw.current))
		bw.current = 0b0
		bw.bitIndex = 0
	}

}

func (bw *BitWriter) YieldSlice() []byte {
	if bw.bitIndex > 0 {
		remaining := 7 - bw.bitIndex
		for range remaining {
			bw.WriteBit(false)
		}
	}
	return bw.storage
}
