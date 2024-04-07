package main

import (
	"fmt"

	fng "github.com/hollowdll/finnish-name-generator"
)

func main() {
	finnishNameGenerator := fng.NewFinnishNameGenerator()

	randomMaleName := finnishNameGenerator.GenerateFirstName(fng.Male)
	randomFemaleName := finnishNameGenerator.GenerateFirstName(fng.Female)
	randomLastName := finnishNameGenerator.GenerateLastName()
	randomGenderFirstName := finnishNameGenerator.GenerateFirstNameRandomGender()

	fmt.Printf("Random male first name: %s\n", randomMaleName)
	fmt.Printf("Random female first name: %s\n", randomFemaleName)
	fmt.Printf("Random last name: %s\n", randomLastName)
	fmt.Printf("Random first name with random gender: %s\n", randomGenderFirstName)
	fmt.Printf("Random full name: %s %s\n", finnishNameGenerator.GenerateFirstNameRandomGender(), finnishNameGenerator.GenerateLastName())
	fmt.Printf("Random male name with two first names: %s %s %s\n", finnishNameGenerator.GenerateFirstName(fng.Male), finnishNameGenerator.GenerateFirstName(fng.Male), finnishNameGenerator.GenerateLastName())
}
