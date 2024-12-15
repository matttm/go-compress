package bitwriter

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
	if set {
		bw.current |= 1
	}
	bw.current << 1
	bw.bitIndex += 1
	if bw.bitIndex == 8 {
		bw.storage = append(bw.storage, byte(bw.current))
		bw.current = 0b0
		bw.bitIndex = 0
	}

}

func (bw *BitWriter) YieldSlice() []byte {
	if bw.bitIndex > 0 {
		remaining := 8 - bw.bitIndex
		for _ = range remaining {
			bw.WriteBit(false)
		}
	}
	return bw.storage
}
