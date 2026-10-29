package main

// there are many crypto/* packages
// hash strength: https://en.wikipedia.org/wiki/Cryptographic_hash_function

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	string_to_hash := "hash me please"
	hash := sha256.New()
	// needs to see strings as a byte slice
	hash.Write([]byte(string_to_hash))

	// use sum to get the finalized hash
	// of you pass a []byte parameter you can append the new hash into an old hash
	hashed := hash.Sum(nil)
	fmt.Println(hashed) // [139 165 136 14 95 168 120 88 44 146 17 48 46 48 159 55 153 219 111 131 72 105 103 187 116 52 54 254 79 51 222 3]
	fmt.Println("String format: ", string(hashed))
	fmt.Printf("Human readable hexadecimal: %x\n", hashed)

}
