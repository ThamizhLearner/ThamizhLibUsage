package main

import (
	"fmt"

	script "github.com/ThamizhLearner/Thamizh"
)

func main() {
	testSubstitution()
	testAppend()
	testTrim()
}

func testSubstitution() {
	s := script.MustDecode("மரங்கள்")

	// Substitution rule definition
	matchTrim := script.MustDecode("ங்கள்")
	substitute := script.MustDecode("ம்")

	s2, ok := s.ReplaceEnd(matchTrim, substitute)
	if !ok {
		fmt.Println("Trim match not found")
		return
	}
	fmt.Printf("Result: %s => %s\n", s, s2)
}

func testAppend() {
	s := script.MustDecode("தமிழ்")
	s2 := script.MustDecode("இலக்கணம்")

	fmt.Printf("%s + %s = %s\n", s, s2, s.Append(s2))
}

func testTrim() {
	s := script.MustDecode("தமிழிலக்கணம்")
	trim := script.MustDecode("இலக்கணம்")

	t, ok := s.TrimEnd(trim)
	if !ok {
		fmt.Println("Trim match not found")
		return
	}
	fmt.Printf("%s - %s = %s\n", s, trim, t)
}
