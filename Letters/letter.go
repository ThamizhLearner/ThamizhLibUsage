package main

import (
	"fmt"

	script "github.com/ThamizhLearner/Thamizh"
)

func main() {
	unicodeStringDecode()

	fmt.Println()
	string2letters("தமிழ்")

	fmt.Println()
	letterAnalysis("தமிழ்")
}

func unicodeStringDecode() {
	fmt.Println("Sample: Decoding unicode string")

	s, ok := script.Decode("தமிழ்")
	if !ok {
		fmt.Println("Decode fail")
		return
	}
	fmt.Println(s)
}

// Get each Thamizh letter from Thamizh string
func string2letters(ustr string) {
	fmt.Println("Sample: Thamizh letters from decode string")

	str := script.MustDecode(ustr)
	fmt.Printf("Thamizh string      : %s\n", str)

	fmt.Println("Thamizh letters forming the Thamizh string:")
	for letter := range str.Letters() {
		fmt.Println(letter)
	}
}

func letterAnalysis(ustr string) {
	fmt.Println("Sample: Split vowelized-consonant into consonant and vowel letters")

	str := script.MustDecode(ustr)
	for letter := range str.Letters() {
		fmt.Print(letter)
		if letter.IsCV() {
			consonant, vowel := letter.SplitCV()
			fmt.Printf(" = %v + %v", consonant, vowel)
		}
		fmt.Println()
	}
}
