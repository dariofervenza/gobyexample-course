package main

// import with alias
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~ TO ENCODE"
	// AGAIN it accepts strings as byte slices
	encoded := b64.StdEncoding.EncodeToString([]byte(data)) // YWJjMTIzIT8kKiYoKSctPUB+IFRPIEVOQ09ERQ==
	fmt.Println("Encoded: ", encoded)

	// decoding, however, may return an err
	decoded, _ := b64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded as str: ", string(decoded))

	// ----- URL COMPATIBLE B64 ENCODING AND DECODING ------
	// ONLY USES url-safe characters in the output
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("URL ENCODING: ", uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
