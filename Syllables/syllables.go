package main

import (
	"fmt"

	script "github.com/ThamizhLearner/Thamizh"
)

func main() {
	s := script.MustDecode("தமிழ்")
	fmt.Println(s)
	fmt.Println(s.SyllabifiedUStr("-"))

	fmt.Println()

	s = script.MustDecode("தமிழிலக்கணம்")
	fmt.Println(s)
	fmt.Println(s.SyllabifiedUStr("-"))

	fmt.Println()

	s = script.MustDecode("காரணமென்னவென்றால்")
	fmt.Println(s)
	fmt.Println(s.SyllabifiedUStr("-"))

	fmt.Println()

	s = script.MustDecode("ஒட்டுமொத்தமாகப்பார்த்துக்கொண்டிருந்தாள்")
	fmt.Println(s)
	fmt.Println(s.SyllabifiedUStr("-"))
}
