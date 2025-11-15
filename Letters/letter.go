package main

import (
	"fmt"

	script "github.com/ThamizhLearner/Thamizh"
)

func main() {
	string2letters("தமிழ்")
}

// Get each Thamizh letter from Thamizh string
func string2letters(ustr string) {
	fmt.Printf("Given Unicode string: %s\n", ustr)

	str := script.MustDecode(ustr)
	fmt.Printf("Thamizh string      : %s\n", str)

	fmt.Println("Thamizh letters forming the Thamizh string:")
	for letter := range str.Values() {
		fmt.Println(letter)
	}
}
