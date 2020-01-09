package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// main compute the hash sha256
func main() {
	arg := os.Args[1]
	h := sha256.New()
	h.Write([]byte(arg))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
