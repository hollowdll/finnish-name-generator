package generator

import (
	"bufio"
	"math/rand"
	"os"
)

type Gender int

const (
	Male Gender = iota
	Female
)

// FinnishNameGenerator is an interface for generating random Finnish first names and last names.
type FinnishNameGenerator interface {
	// GenerateFirstName generates a first name based on gender.
	GenerateFirstName(gender Gender) string
	// GenerateFirstNameRandomGender generates a random gender first name.
	GenerateFirstNameRandomGender() string
	// GenerateLastName generates a last name.
	GenerateLastName() string
}

// finnishNameGenerator is a struct that implements FinnishNameGenerator.
type finnishNameGenerator struct{}

// NewFinnishNameGenerator returns a new FinnishNameGenerator, or error if failed.
func NewFinnishNameGenerator() FinnishNameGenerator {
	return &finnishNameGenerator{}
}

func (f *finnishNameGenerator) GenerateFirstName(gender Gender) string {
	switch gender {
	case Male:
		return maleFirstNames[rand.Intn(len(maleFirstNames))]
	case Female:
		return femaleFirstNames[rand.Intn(len(femaleFirstNames))]
	default:
		return ""
	}
}

func (f *finnishNameGenerator) GenerateFirstNameRandomGender() string {
	rngGender := rand.Intn(2)

	if rngGender == 0 {
		return maleFirstNames[rand.Intn(len(maleFirstNames))]
	} else {
		return femaleFirstNames[rand.Intn(len(femaleFirstNames))]
	}
}

func (f *finnishNameGenerator) GenerateLastName() string {
	return lastNames[rand.Intn(len(lastNames))]
}

// readFileLines reads the lines of filepath and returns them as a string slice.
func readFileLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
