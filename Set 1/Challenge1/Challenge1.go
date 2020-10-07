package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

//Challenge1 of the set
func main() {
	// did the reverse of hex to base64
	docID := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	base64Decode, err := base64.StdEncoding.DecodeString(docID)
	if err != nil {
		log.Fatal("error:", err)
	}
	base64Decoded := fmt.Sprintf("%q", base64Decode)
	fmt.Printf("base_decoded %v\n", base64Decoded)

	src := []byte(base64Decoded)
	hexEncode := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(hexEncode, src)
	hexEncoded := fmt.Sprintf("%s", hexEncode)
	fmt.Printf("hex_encoded %v", hexEncoded)
}
