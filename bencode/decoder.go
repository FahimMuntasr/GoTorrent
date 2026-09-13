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
	// Check for out of bounds 
	if d.pos >= len(d.data) {
		return 0, fmt.Errorf("out of bounds index")
	}
	
	// Check for missing i prefix 
	if d.data[d.pos] != 'i' {
		return 0, fmt.Errorf("expected integer")
	}

	// initialize start of int 
	d.pos++
	start := d.pos
	
	// loop through byte array
	for d.pos < len(d.data) && d.data[d.pos] != 'e' {
		d.pos++
	}
	
	// Check for missing e suffix
	if d.pos >= len(d.data) {
    return 0, fmt.Errorf("unterminated integer")
	}
 
	// Check negative zero and leading zero 
	if d.pos - start > 1 {
		if d.data[start] == '-' && d.data[start + 1] == '0' {
			return 0, fmt.Errorf("negative integer has leading zero")
		}

		if d.data[start] == '0' {
			return 0, fmt.Errorf("leading zero")
		}
	}

	// Extract and convert string to 64 bit int 
	numberString := string(d.data[start:d.pos])

	numberInt, err := strconv.ParseInt(numberString, 10, 64);
	
	if err != nil {
		return 0, err
	}

	// move position after end of current int 
	d.pos++
	
	return numberInt, nil 
}

func (d *Decoder) decodeString() (string, error) {
	// Initialize start and loop through till limiter 
	start := d.pos

	for d.pos < len(d.data) && d.data[d.pos] != ':' {
		d.pos++
	}
	
	// Check if limiter found 
	if d.pos >= len(d.data) {
		return "", fmt.Errorf("Missing string delimiter")
	}
	
	// extract and convert string to int 
	length, err := strconv.Atoi(string(d.data[start:d.pos]))

	if err != nil {
		return "", err 
	}
	
	// make sure size is accurate 
	d.pos++

	if d.pos + length > len(d.data) {
		return "", fmt.Errorf("string length not accurate")
	}

	// extract string 
	result := string(d.data[d.pos : d.pos+length])

	// increase position to end and return string 
	
	d.pos += length 

	return result, nil
}

func (d *Decoder) decodeList() ([]any, error) {
  if d.pos >= len(d.data) {
    return nil, fmt.Errorf("out of bounds index")
  }
	
	// check that current byte is 'l'
	if d.data[d.pos] != 'l'{
		return nil, fmt.Errorf("expected list")
	}
	list := []any{}
  
	// move past 'l'
	d.pos++

  // keep decoding items until we encounter 'e'
	for d.pos < len(d.data) && d.data[d.pos] != 'e' {
		item, err := d.Decode()
		if err != nil {
    	return nil, err
		}
		list = append(list, item)
	}
	if d.pos >= len(d.data) {
    return nil, fmt.Errorf("unterminated list")
	}

  // move past 'e'
	d.pos++
  // return the list
	return list, nil 
}

func (d *Decoder) decodeDictionary() (map[string]any, error) {
  if d.pos >= len(d.data) {
    return nil, fmt.Errorf("out of bounds index")
  }
	
	if d.data[d.pos] != 'd'{
		return nil, fmt.Errorf("expected dictionary")
	}
  // Move past 'd'
	d.pos++
  // Create dictionary
	dict := map[string]any{}
	
  // Loop until 'e'
	for d.pos < len(d.data) && d.data[d.pos] != 'e' {
    var key string
		var value any
		var err error
		// Decode key
		key, err = d.decodeString()
		if err != nil {
			return nil, err
		}
    // Decode value
		value, err = d.Decode()
		if err != nil {
    	return nil, err
		}
    // Add key/value to dictionary
		dict[key] = value
	}

  // Check for terminating 'e'
	if d.pos >= len(d.data) {
		return nil, fmt.Errorf("unterminated dictionary")
	}

  // Move past 'e'
	d.pos++

  // Return dictionary
	return dict, nil 
}

func (d *Decoder) Decode() (any, error) {
    if d.pos >= len(d.data) {
        return nil, fmt.Errorf("out of bounds index")
    }

    switch d.data[d.pos] {
    case 'i':
        return d.decodeInteger()

    case 'l':
        return d.decodeList()

    case 'd':
        return d.decodeDictionary()

    default:
        return d.decodeString()
    }
}
