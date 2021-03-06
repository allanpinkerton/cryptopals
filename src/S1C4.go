package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func getEngishFrequencyMap() map[rune]float32 {
	var frequencyMap = make(map[rune]float32)
	frequencyMap['e'] = 12.702
	frequencyMap['t'] = 9.356
	frequencyMap['a'] = 8.167
	frequencyMap['o'] = 7.507
	frequencyMap['i'] = 6.966
	frequencyMap['n'] = 6.749
	frequencyMap['s'] = 6.327
	frequencyMap['h'] = 6.094
	frequencyMap['r'] = 5.987
	frequencyMap['d'] = 4.253
	frequencyMap['l'] = 4.025
	frequencyMap['u'] = 2.758
	frequencyMap['w'] = 2.560
	frequencyMap['m'] = 2.406
	frequencyMap['f'] = 2.228
	frequencyMap['c'] = 2.202
	frequencyMap['g'] = 2.015
	frequencyMap['y'] = 1.994
	frequencyMap['p'] = 1.929
	frequencyMap['b'] = 1.492
	frequencyMap['k'] = 1.292
	frequencyMap['v'] = 0.978
	frequencyMap['j'] = 0.153
	frequencyMap['x'] = 0.150
	frequencyMap['q'] = 0.095
	frequencyMap['z'] = 0.077

	return frequencyMap
}

func score(s string, frequencyMap map[rune]float32) float32 {
	var total float32 = 0
	for _, char := range s {
		total += frequencyMap[char]
	}
	return total
}

func guessEncrypted(inputHex string, index int) {
	frequencyMap := getEngishFrequencyMap()

	var maxScore float32 = 0
	for i := 0; i < 256; i++ {
		bsIn, _ := hex.DecodeString(inputHex)
		for j, _ := range bsIn {
			bsIn[j] = bsIn[j] ^ byte(i)
		}
		s := score(string(bsIn), frequencyMap)
		if s > maxScore {
			maxScore = s
		}
		if s > 100.0 {
			fmt.Printf("Line %d Score: %f\n", index, s)
			fmt.Printf("Byte being XORED%x\n", i)
			fmt.Println("Guessed text is:", string(bsIn))
			fmt.Println("=========================")
		}
	}

}

func readDataIntoArray(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	lines := readDataIntoArray("/Users/jma/dev/cryptopals/assets/S1C4.data")
	for i, _ := range lines {
		guessEncrypted(lines[i], i)
	}
}
