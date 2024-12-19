package bitreader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type bool_testcase struct {
	input    []byte
	expected []bool
}

func Test_ShouldReadSingleBits(t *testing.T) {
	tests := []bool_testcase{
		{
			input:    []byte{0b00000001},
			expected: []bool{false, false, false, false, false, false, false, true},
		},
		{
			input:    []byte{0b10000001},
			expected: []bool{true, false, false, false, false, false, false, true},
		},
		{
			input:    []byte{0b11000001},
			expected: []bool{true, true, false, false, false, false, false, true},
		},
		{
			input:    []byte{0b11000001, 0b11000001},
			expected: []bool{true, true, false, false, false, false, false, true, true, true, false, false, false, false, false, true},
		},
	}
	for _, _t := range tests {
		storage := _t.input
		bw := FromSlice(storage)
		for i, _ := range _t.expected {
			set, _ := bw.ReadBit()
			assert.Equal(t, set, _t.expected[i], fmt.Sprint(i))
		}
	}
}
