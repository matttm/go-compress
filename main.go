package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/matttm/go-compress/internal/huffman"
)

// Usage: cat sample | ./gocompress
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			continue
		}
		sb.WriteString(line)
	}
	huffman.FromReader(false, sb.String())
}
