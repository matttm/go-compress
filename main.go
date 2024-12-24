package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/matttm/go-compress/internal/huffman"
	"github.com/urfave/cli/v2"
)

// program for encoding/decoding data
//
// Usage: cat sample | ./gocompress encode

// function pipes all stdin as input to go-compress
func PipeStdInToCommand(fn func(string) *huffman.HuffmanCodec) {
	scanner := bufio.NewScanner(os.Stdin)
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			continue
		}
		sb.WriteString(line)
	}
	fn(sb.String())
}
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "encode",
				Aliases: []string{},
				Usage:   "encode <content stream>",
				Action: func(cCtx *cli.Context) error {
					// path := cCtx.Args().Get(0)
					PipeStdInToCommand(huffman.FromDecodedText)
					return nil
				},
			},
			{
				Name:    "decode",
				Aliases: []string{},
				Usage:   "encode <content stream>",
				Action: func(cCtx *cli.Context) error {
					// path := cCtx.Args().Get(0)
					PipeStdInToCommand(huffman.FromEncodedText)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
