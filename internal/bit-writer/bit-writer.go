package bitwriter

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

}

// yields its internal buffer to a slice and a integer
//
//	indicating how many bits were used to fill the last byte
//
// NOTE: these remaining bit count should be stored with encoded data,
//
//	so these extera bits can be ignored during deserialization
func (bw *BitWriter) YieldSlice() ([]byte, uint8) {
	var remaining uint8 = 0
	if bw.bitCount > 0 {
		remaining = 8 - bw.bitCount
		for range remaining {
			bw.WriteBit(false)
		}
	}
	return bw.storage, remaining
}
