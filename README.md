# go-compress

## Description

This codec written in go, currently encodes/decodes using the Huffman algorithm.

## Getting Started
There are two commands `encode` and `decode`, for which I normally pipe in input. This might look like the following:
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
## Design

The following is a quick diagram, of the encoded data
```
+-------------------------------------------------+
|                  Header (2 Bytes)               | 
+-------------------------------------------------+
|              Extraneous Bits (1 Byte)           | 
+-------------------------------------------------+
|              Serialized Tree (k Bytes)          | 
+-------------------------------------------------+
|                 Encoded Data                    | 
+-------------------------------------------------+
```
The header is simply two bytes indicating the following is encoded data.

The following byte is an integer in the range from zero to seven, indicating how many bits should be ignored in the last byte of data.

The following bytes are of a variable amount. However though the length is unknown, since null values are presented after leaves, the boundary can be determined.

The rest of the bytes are encoded data.

For a better understanding of the extraneous bits byte, if the byte is `0x03`, then of the last byte of encoded data, the last theee bits are only for byte alignment and to ensure the decode does not determine these to be a character in the plaintext.

## Authors

-   Matt Maloney : matttm

## Contribute

If you want to contribute, just send me a message.
