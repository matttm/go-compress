package bitwriter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type bool_testcase struct {
	input    []bool
	expected []byte
}

func Test_ShouldEncode(t *testing.T) {
	tests := []bool_testcase{
		{
			input:    []bool{true},
			expected: []byte{0b10000000},
		},
		{
			input:    []bool{true, false, false, true},
			expected: []byte{0b10010000},
		},
		{
			input:    []bool{true, false, false, true, true, false, false, true},
			expected: []byte{0b10011001},
		},
	}
	for _, _t := range tests {
		storage := []byte{}
		bw := WithSlice(storage)
		for _, _bool := range _t.input {
			bw.WriteBit(_bool)
		}
		yield := bw.YieldSlice()
		assert.Equal(t, len(yield), 1, "")
		assert.Equal(t, yield[0], _t.expected[0], "should pad to byte boundary")
	}
}
