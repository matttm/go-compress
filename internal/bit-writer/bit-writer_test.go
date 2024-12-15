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
	}
	for _, _t := range tests {
		storage := []byte{}
		bw := WithSlice(storage)
		for _, _bool := range _t.input {
			bw.WriteBit(_bool)
		}
		yield := bw.YieldSlice()
		assert.Equal(t, yield[0], _t.expected[0], "swcjwdin")
	}
}
