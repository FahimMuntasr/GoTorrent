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

func TestDecodeString(t *testing.T) {
  decoder := NewDecoder([]byte("4:spam"))

  str, err := decoder.decodeString()

  if err != nil {
      t.Fatal(err)
  }

  if str != "spam" {
      t.Fatalf("expected spam, got %s", str)
  }
}

func TestDecodeEmptyString(t *testing.T) {
  decoder := NewDecoder([]byte("0:"))

  str, err := decoder.decodeString()

  if err != nil {
      t.Fatal(err)
  }

  if str != "" {
      t.Fatalf("expected empty string, got %s", str)
  }
}

func TestDecodeStringWithRemainingData(t *testing.T) {
    decoder := NewDecoder([]byte("4:spami42e"))

    str, err := decoder.decodeString()

    if err != nil {
        t.Fatal(err)
    }

    if str != "spam" {
        t.Fatalf("expected spam, got %s", str)
    }

    if decoder.pos != 6 {
        t.Fatalf("expected position 6, got %d", decoder.pos)
    }
}

func TestDecodeStringNoLength(t *testing.T) {
	decoder := NewDecoder([]byte(":e"))

	_, err := decoder.decodeString()

	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestDecodeIncompleteString(t *testing.T) {
	decoder := NewDecoder([]byte("5:spam"))

	_, err := decoder.decodeString()

	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestDecodeStringLengthOverflow(t *testing.T) {
	decoder := NewDecoder([]byte("999999999999999999999999:spam"))

	_, err := decoder.decodeString()

	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestDecodeList(t *testing.T) {
    decoder := NewDecoder([]byte("l4:spam4:eggse"))

    list, err := decoder.decodeList()

    if err != nil {
        t.Fatal(err)
    }

    if len(list) != 2 {
        t.Fatalf("expected 2 items, got %d", len(list))
    }

    if list[0] != "spam" {
        t.Fatalf("expected spam, got %v", list[0])
    }

    if list[1] != "eggs" {
        t.Fatalf("expected eggs, got %v", list[1])
    }
}

func TestDecodeMixedList(t *testing.T) {
    decoder := NewDecoder([]byte("li42e4:spame"))

    list, err := decoder.decodeList()

    if err != nil {
        t.Fatal(err)
    }

    if len(list) != 2 {
        t.Fatalf("expected 2 items, got %d", len(list))
    }

    if list[0] != int64(42) {
        t.Fatalf("expected 42, got %v", list[0])
    }

    if list[1] != "spam" {
        t.Fatalf("expected spam, got %v", list[1])
    }
}

func TestDecodeEmptyList(t *testing.T) {
    decoder := NewDecoder([]byte("le"))

    list, err := decoder.decodeList()

    if err != nil {
        t.Fatal(err)
    }

    if len(list) != 0 {
        t.Fatalf("expected empty list, got %d items", len(list))
    }
}

func TestDecodeNestedList(t *testing.T) {
    decoder := NewDecoder([]byte("lli1ei2eee"))

    list, err := decoder.decodeList()

    if err != nil {
        t.Fatal(err)
    }

    if len(list) != 1 {
        t.Fatalf("expected 1 item, got %d", len(list))
    }

    nested, ok := list[0].([]any)
    if !ok {
        t.Fatalf("expected nested list, got %T", list[0])
    }

    if len(nested) != 2 {
        t.Fatalf("expected 2 items in nested list, got %d", len(nested))
    }

    if nested[0] != int64(1) {
        t.Fatalf("expected 1, got %v", nested[0])
    }

    if nested[1] != int64(2) {
        t.Fatalf("expected 2, got %v", nested[1])
    }
}

func TestDecodeUnterminatedList(t *testing.T) {
    decoder := NewDecoder([]byte("l4:spam"))

    _, err := decoder.decodeList()

    if err == nil {
        t.Fatal("expected an error")
    }
}

func TestDecodeDictionary(t *testing.T) {
	decoder := NewDecoder([]byte("d3:cow3:mooe"))

	dict, err := decoder.decodeDictionary()

	if err != nil {
		t.Fatal(err)
	}

	if dict["cow"] != "moo" {
		t.Fatalf("expected moo, got %v", dict["cow"])
	}
}

func TestDecodeEmptyDictionary(t *testing.T) {
	decoder := NewDecoder([]byte("de"))

	dict, err := decoder.decodeDictionary()

	if err != nil {
		t.Fatal(err)
	}

	if len(dict) != 0 {
		t.Fatalf("expected empty dictionary, got %d items", len(dict))
	}
}

func TestDecodeMultipleDictionaryItems(t *testing.T) {
	decoder := NewDecoder([]byte("d3:cow3:moo4:spam4:eggse"))

	dict, err := decoder.decodeDictionary()

	if err != nil {
		t.Fatal(err)
	}

	if dict["cow"] != "moo" {
		t.Fatalf("expected moo, got %v", dict["cow"])
	}

	if dict["spam"] != "eggs" {
		t.Fatalf("expected eggs, got %v", dict["spam"])
	}
}

func TestDecodeMixedDictionary(t *testing.T) {
	decoder := NewDecoder([]byte("d3:agei20e4:name5:Fahime"))

	dict, err := decoder.decodeDictionary()

	if err != nil {
		t.Fatal(err)
	}

	if dict["age"] != int64(20) {
		t.Fatalf("expected 20, got %v", dict["age"])
	}

	if dict["name"] != "Fahim" {
		t.Fatalf("expected Fahim, got %v", dict["name"])
	}
}

func TestDecode(t *testing.T) {
    tests := []struct {
        name string
        data string
        want any
    }{
        {
            name: "integer",
            data: "i42e",
            want: int64(42),
        },
        {
            name: "string",
            data: "4:spam",
            want: "spam",
        },
        {
            name: "list",
            data: "li42e4:spame",
            want: []any{int64(42), "spam"},
        },
        {
            name: "dictionary",
            data: "d3:agei20e4:name5:Fahime",
            want: map[string]any{
                "age":  int64(20),
                "name": "Fahim",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            decoder := NewDecoder([]byte(tt.data))

            got, err := decoder.Decode()
            if err != nil {
                t.Fatal(err)
            }

            // We'll improve this comparison shortly.
            _ = got
        })
    }
}













