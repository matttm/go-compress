# go-compress

## Description

This codec written in go, currently encodes/decodes using the Huffman algorithm.

## Getting Started
There are two commands `encode` and `decode`, for which I normally pipe in input. This might look the following:
```
❯ echo "aabba" | ./go-compress encode
*b@@a@@
```
and decoding could then look like:
```
❯ echo "aabba" | ./go-compress encode | ./go-compress decode
aabba
```
so if you want to encode a file, use some combo of the `cat` command and file redirection, something like:
```
cat decoded.file | ./go-compress encode > encoded.file
```

## Authors

-   Matt Maloney : matttm

## Contribute

If you want to contribute, just send me a message.
