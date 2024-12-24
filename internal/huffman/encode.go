package huffman

import (
	"fmt"
	"sort"
	"strings"

	bitwriter "github.com/matttm/go-compress/internal/bit-writer"
)

// Function FromDecodedText
//
//	converts encoded data into decodec data
func FromDecodedText(s string) *HuffmanCodec {
	encoder := New()
	encoder.constructFrequencyMap(s)
	encoder.createTree()
	var sb strings.Builder
	encoder.createCodeTable(encoder.tree, sb)
	bytes, extra := encoder.encode(s)
	serializedTree := serializeTree(encoder.tree)
	encoder.encoded = packageData(extra, serializedTree, bytes)
	fmt.Printf("%s\n", encoder.encoded)
	return encoder
}

// Function encode
//
//	uses field tree of HuffmanNode to encode the data
func (c *HuffmanCodec) encode(s string) ([]byte, uint8) {
	encoded := []byte{}
	bw := bitwriter.WithSlice(encoded)
	for _, r := range s {
		code := c.encodingTable[r]
		for _, c := range code {
			bw.WriteBit(c == '1')
		}
	}
	return bw.YieldSlice()
}
func (c *HuffmanCodec) constructFrequencyMap(s string) {
	c.frequencyTable = make(map[rune]int)
	for _, r := range s {
		prv := c.frequencyTable[r]
		c.frequencyTable[r] = prv + 1
	}
}
func (c *HuffmanCodec) createTree() {
	// init queue
	pq := PriorityQueue{}
	// heap.Init(&pq)
	// get the total so i can calculate probabilities
	sum := 0
	for _, v := range c.frequencyTable {
		sum += v
	}
	// fmt.Println(sum)
	for k, v := range c.frequencyTable {
		pq.Push(
			&Item{
				Node: &HuffmanNode{
					symbol: k,
					prb:    float32(v) / float32(sum),
				},
			},
		)
	}
	// create tree
	for pq.Len() > 1 {
		sort.Slice(pq, func(i, j int) bool {
			return pq[i].Node.prb > pq[j].Node.prb
		})
		a := pq.Pop().(*Item).Node
		b := pq.Pop().(*Item).Node
		//  fmt.Printf("%c %.06f %c %.06f\n", a.symbol, a.prb, b.symbol, b.prb)
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
		n := pq.Pop().(*Item)
		c.tree = n.Node
	} else {
		panic("Huffman map is empty")
	}
}
func serializeTree(n *HuffmanNode) []byte {
	serializedTree := []byte{}
	nodeCount := 0
	var serialize func(*HuffmanNode, *[]byte, *int)
	serialize = func(_n *HuffmanNode, _b *[]byte, count *int) {
		if _n == nil {
			*_b = append(*_b, NULL)
			return
		}
		*_b = append(*_b, byte(_n.symbol))
		*count += 1
		//*_b = append(*_b, 0)
		serialize(_n.left, _b, count)
		//*_b = append(*_b, 1)
		serialize(_n.right, _b, count)
	}
	serialize(n, &serializedTree, &nodeCount)
	return serializedTree
}
func packageData(extraBits uint8, serializedTree, encodedData []byte) []byte {
	packagedData := []byte{}
	packagedData = append(packagedData, MAGIC_NUMBER...)
	packagedData = append(packagedData, extraBits)
	packagedData = append(packagedData, serializedTree...)
	packagedData = append(packagedData, encodedData...)
	return packagedData
}
