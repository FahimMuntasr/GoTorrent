package bencode

type Decoder struct {
	data []byte
	pos int
}

func NewDecoder(data []byte) *Decoder {
	return &Decoder {
		data: data,
		pos: 0,
	}
}
