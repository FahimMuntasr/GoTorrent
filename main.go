package main

import (
	"fmt"

	"github.com/FahimMuntasr/GoTorrent/bencode"
)

func main() {
	decoder := bencode.NewDecoder([]byte("i42e"))

	fmt.Println(decoder)
}
