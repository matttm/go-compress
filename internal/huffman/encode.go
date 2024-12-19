package huffman

import (
	"fmt"
	"sort"
	"strings"

	bitwriter "github.com/matttm/go-compress/internal/bit-writer"
)

func (c *HuffmanCodec) encode(s string) []byte {
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

func FromDecodedText(s string) *HuffmanCodec {
	encoder := new(HuffmanCodec)
	encoder.constructFrequencyMap(s)
	encoder.createTree()
	var sb strings.Builder
	createCodeTable(encoder.tree, encoder.encodingTable, sb)
	bytes := encoder.encode(s)
	encoder.encoded = bytes
	//	fmt.Println(encoder.encodingTable)
	//	fmt.Printf("%08b\n", bytes)
	return encoder
}
func (c *HuffmanCodec) constructFrequencyMap(s string) {
	c.frequencyTable = make(map[rune]int)
	c.encodingTable = make(map[rune]string)
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
	fmt.Println(sum)
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
func createCodeTable(node *HuffmanNode, m map[rune]string, sb strings.Builder) {
	if node == nil {
		return
	}
	if node.symbol != '*' {
		m[node.symbol] = sb.String()
		return
	}
	if node.left != nil {
		var left strings.Builder
		left.WriteString(sb.String())
		left.WriteRune('0')
		createCodeTable(node.left, m, left)
	}
	if node.right != nil {
		var right strings.Builder
		right.WriteString(sb.String())
		right.WriteRune('1')
		createCodeTable(node.right, m, right)
	}
}
func serializeTree(n *HuffmanNode) []byte {
	b := []byte{}
	nodeCount := 0
	var serialize func(*HuffmanNode, *[]byte, *int)
	serialize = func(_n *HuffmanNode, _b *[]byte, count *int) {
		if _n == nil {
			return
		}
		if _n.symbol != '*' {
			*_b = append(*_b, LEAF...)
			*_b = append(*_b, byte(_n.symbol))
			*count += 1
		} else {
			*_b = append(*_b, INTER...)
			serialize(_n.left, _b, count)
			serialize(_n.left, _b, count)
		}
	}
	serialize(n, &b, &nodeCount)
	encodedData := []byte{}
	encodedData = append(encodedData, MAGIC_NUMBER...)
	encodedData = append(encodedData, b...)
	return encodedData
}
