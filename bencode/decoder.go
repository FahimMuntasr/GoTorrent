package bencode
import "strconv"
import "fmt"

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

func (d *Decoder) decodeInteger() (int64, error) {
	if d.pos >= len(d.data) {
		return 0, fmt.Errorf("out of bounds index")
	}

	if d.data[d.pos] != 'i' {
		return 0, fmt.Errorf("expected integer")
	}

	d.pos++
	start := d.pos
	
	for d.pos < len(d.data) && d.data[d.pos] != 'e' {
		d.pos++
	}
	
	if d.pos >= len(d.data) {
    return 0, fmt.Errorf("unterminated integer")
	}
 
	// negative zero 
	if d.pos - start > 1 {
		if d.data[start] == '-' && d.data[start + 1] == '0' {
			return 0, fmt.Errorf("negative integer has leading zero")
		}

		if d.data[start] == '0' {
			return 0, fmt.Errorf("leading zero")
		}
	}

	
	numberString := string(d.data[start:d.pos])

	numberInt, err := strconv.ParseInt(numberString, 10, 64);
	
	if err != nil {
		return 0, err
	}

	d.pos++
	
	return numberInt, nil 
}
