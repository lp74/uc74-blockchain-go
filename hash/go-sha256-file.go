package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// main compute the hash sha256 of a file
func main() {
	arg := os.Args[1]
	fmt.Println(arg)

	dat, err := ioutil.ReadFile(arg)
	check(err)
	fmt.Print(string(dat))
	h := sha256.New()
	h.Write([]byte(dat))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
