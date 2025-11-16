package main

import (
	"fmt"

	script "github.com/ThamizhLearner/Thamizh"
)

func main() {
	ustr := "தமிழ்"
	fmt.Printf("Given Unicode string: %s\n", ustr)
	string2letters(ustr)

	fmt.Println("\nLetter split")
	letterAnalysis(ustr)
}

// Get each Thamizh letter from Thamizh string
func string2letters(ustr string) {

	str := script.MustDecode(ustr)
	fmt.Printf("Thamizh string      : %s\n", str)

	fmt.Println("Thamizh letters forming the Thamizh string:")
	for letter := range str.Values() {
		fmt.Println(letter)
	}
}

func letterAnalysis(ustr string) {
	str := script.MustDecode(ustr)

	for letter := range str.Values() {
		fmt.Print(letter)
		if letter.IsCV() {
			consonant, vowel := letter.DetachedCV()
			fmt.Printf(" = %v + %v", consonant, vowel)
		}
		fmt.Println()
	}
}
