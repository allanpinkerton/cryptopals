package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	byteArr, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	b64String := base64.StdEncoding.EncodeToString(byteArr)
	fmt.Println(b64String)
}
