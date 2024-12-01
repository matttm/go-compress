package gocompress

import (
	"container/heap"
	"io"
)

type HuffmanCodec struct {
	frqMap map[rune]int
	tree   *HuffmanNode
}

func FromReader(isCompressed bool, r io.Reader) *HuffmanCodec {
	encoder := new(HuffmanCodec)
	if isCompressed {
	} else {
		encoder.constructFrequencyMap(r)
	}
	return encoder
}
func (c *HuffmanCodec) constructFrequencyMap(r io.Reader) {
	c.frqMap = make(map[rune]int)
	var b []byte
	_, err := r.Read(b)
	if err != nil {
		panic(err)
	}
	for _, r := range string(b) {
		prv := c.frqMap[r]
		c.frqMap[r] = prv + 1
	}
}
func (c *HuffmanCodec) createTree() {
	// init queue
	pq := PriorityQueue{}
	heap.Init(&pq)
	// get the total so i can calculate probabilities
	sum := 0
	for _, v := range c.frqMap {
		sum += v
	}
	for k, v := range c.frqMap {
		pq.Push(
			HuffmanNode{
				symbol: k,
				prb:    float32(v) / float32(sum),
			},
		)
	}
	// create tree
	for pq.Len() > 1 {
		a := pq.Pop().(Item).Node
		b := pq.Pop().(Item).Node
		pq.Push(
			&Item{
				Node: &HuffmanNode{
					symbol: '*',
					prb:    a.prb + b.prb,
					left:   a,
					right:  b,
				},
			},
		)
	}
	if pq.Len() == 1 {
		n := pq.Pop().(Item)
		c.tree = n.Node
	} else {
		panic("Huffman hdap is empty")
	}
}
