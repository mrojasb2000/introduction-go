package main

import (
	"fmt"
	"hash/crc32"
)

// Non-cryptographic
func main() {
	// create a hasher
	h := crc32.NewIEEE()
	// write out data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

}
