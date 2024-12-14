package codec

import "io"

type Codec interface {
	Encode(w io.WriteSeeker)
	Decode(w io.WriteSeeker)
}
