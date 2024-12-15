package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type huffman_testcase struct {
	input    string
	expected []byte
}

func Test_ShouldEncode(t *testing.T) {
	tests := []huffman_testcase{}
	for _, _t := range tests {
		c := FromDecodedText(false, _t.input)
		assert.ElementsMatch(t, c.encoded, _t.expected, "swcjwdin")
	}
}
