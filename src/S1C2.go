package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"
	arr1, _ := hex.DecodeString(s1)
	arr2, _ := hex.DecodeString(s2)
	for i, _ := range arr1 {
		arr1[i] = arr1[i] ^ arr2[i]
	}
	result := hex.EncodeToString(arr1)

	fmt.Println(result)
	fmt.Println("ASCII is:", string(arr1))
}
