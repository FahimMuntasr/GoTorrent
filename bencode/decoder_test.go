package bencode

import "testing"

func TestDecodeInteger(t *testing.T) {
	decoder := NewDecoder([]byte("i-42e"))

	integer, err := decoder.decodeInteger()
	
	if err != nil {
		t.Fatal(err)
	}
	
	if integer != -42 {
		t.Fatalf("expected -42, got %d", integer)
	}
}

func TestDecodeUnterminatedInteger(t *testing.T) {
  decoder := NewDecoder([]byte("i42"))

  _, err := decoder.decodeInteger()

  if err == nil {
      t.Fatal("expected an error")
  }
}

func TestDecodeNoNumberInteger(t *testing.T) {
  decoder := NewDecoder([]byte("ie"))

  _, err := decoder.decodeInteger()

  if err == nil {
      t.Fatal("expected an error")
  }
}

func TestDecodeNegativeZeroInteger(t *testing.T) {
  decoder := NewDecoder([]byte("i-0e"))

  _, err := decoder.decodeInteger()

  if err == nil {
      t.Fatal("expected an error")
  }
}

func TestDecodeLeadingZeroInteger(t *testing.T) {
  decoder := NewDecoder([]byte("i042e"))

  _, err := decoder.decodeInteger()

  if err == nil {
      t.Fatal("expected an error")
  }
}

func TestDecodeNegativeLeadingZeroInteger(t *testing.T) {
  decoder := NewDecoder([]byte("i-042e"))

  _, err := decoder.decodeInteger()

  if err == nil {
      t.Fatal("expected an error")
  }
}

func TestDecodeNonInteger(t *testing.T) {
	decoder := NewDecoder([]byte("42e"))

	_, err := decoder.decodeInteger()

	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestOverflowInteger(t *testing.T) {
	decoder := NewDecoder([]byte("i999999999999999999999999999e"))

	_, err := decoder.decodeInteger()

	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestInvalidInteger(t *testing.T) {
	decoder := NewDecoder([]byte("i-e"))

	_, err := decoder.decodeInteger()

	if err == nil {
		t.Fatal("expected an error")
	}
}
