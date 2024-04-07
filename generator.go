package generator

import (
	"bufio"
	"math/rand"
	"os"
	"time"
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
type finnishNameGenerator struct {
	maleFirstNames   []string
	femaleFirstNames []string
	lastNames        []string
}

// NewFinnishNameGenerator returns a new FinnishNameGenerator, or error if failed.
func NewFinnishNameGenerator() (FinnishNameGenerator, error) {
	maleFirstNames, err := readFileLines("dataset/testdata/male_first_names.txt")
	if err != nil {
		return nil, err
	}

	femaleFirstNames, err := readFileLines("dataset/testdata/female_first_names.txt")
	if err != nil {
		return nil, err
	}

	lastNames, err := readFileLines("dataset/testdata/last_names.txt")
	if err != nil {
		return nil, err
	}

	return &finnishNameGenerator{
		maleFirstNames:   maleFirstNames,
		femaleFirstNames: femaleFirstNames,
		lastNames:        lastNames,
	}, nil
}

func (f *finnishNameGenerator) GenerateFirstName(gender Gender) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch gender {
	case Male:
		return f.maleFirstNames[rng.Intn(len(f.maleFirstNames))]
	case Female:
		return f.femaleFirstNames[rng.Intn(len(f.femaleFirstNames))]
	default:
		return ""
	}
}

func (f *finnishNameGenerator) GenerateFirstNameRandomGender() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rngGender := rng.Intn(2)

	if rngGender == 0 {
		return f.maleFirstNames[rng.Intn(len(f.maleFirstNames))]
	} else {
		return f.femaleFirstNames[rng.Intn(len(f.femaleFirstNames))]
	}
}

func (f *finnishNameGenerator) GenerateLastName() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	return f.lastNames[rng.Intn(len(f.lastNames))]
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
