package main

import (
	"bufio"
	"os"
	"strings"
)

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
	FromReader(false, sb.String())
}
