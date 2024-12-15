package huffman

import (
	"fmt"
	"sort"
	"strings"
)

type HuffmanCodec struct {
	frequencyTable map[rune]int
	codeTable      map[rune]string
	tree           *HuffmanNode
}

func (c *HuffmanCodec) encode(s string) []byte {
	encoded := []byte{}
	for _, r := range s {
		var b byte = 0
		var bitPos uint8 = 7
		code := c.codeTable[r]
		for _, c := range code {
			if c == '1' {
				// set bit
				b |= 1 << bitPos
			} else {
				// unset bit
				b &= ^(1 << bitPos)
			}
			bitPos -= 1
			resetBytePos(&bitPos, b, encoded)
		}
	}
	return encoded
}

func FromDecodedText(isCompressed bool, s string) *HuffmanCodec {
	encoder := new(HuffmanCodec)
	fmt.Printf("Huffman input: %s", s)
	if isCompressed {
	} else {
		encoder.constructFrequencyMap(s)
		fmt.Println(encoder.frequencyTable)
		encoder.createTree()
		var sb strings.Builder
		createCodeTable(encoder.tree, encoder.codeTable, sb)
		fmt.Println(encoder.codeTable)
	}
	bytes := encoder.encode(s)
	fmt.Println(bytes)
	return encoder
}
func (c *HuffmanCodec) constructFrequencyMap(s string) {
	c.frequencyTable = make(map[rune]int)
	c.codeTable = make(map[rune]string)
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

func resetBytePos(bitPos *uint8, b uint8, encoded []byte) {
	if *bitPos > 8 {
		*bitPos = 7
		encoded = append(encoded, b)
	}
}
