package main

import (
	"fmt"
	"strconv"
	"strings"

	"../utils"
)

const invalidValue = -1

type Aunt struct {
	id int

	children int

	cats int

	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int

	goldfish int

	trees int

	cars int

	perfumes int
}

func (a *Aunt) Parse(line string) {
	tokens := strings.Split(line, " ")
	if len(tokens) <= 2 {
		panic(tokens)
	}
	if tokens[0] != "Sue" {
		panic(tokens)
	}
	a.id, _ = strconv.Atoi(utils.ReadUntilLast(tokens[1]))

	// TODO: use reflection
	var value int = invalidValue
	a.children = value
	a.cats = value
	a.samoyeds = value
	a.pomeranians = value
	a.akitas = value
	a.vizslas = value
	a.goldfish = value
	a.trees = value
	a.cars = value
	a.perfumes = value

	if len(tokens[2:])%2 != 0 {
		panic(tokens)
	}
	for i := 1; i < len(tokens)/2; i++ {
		key := utils.ReadUntilLast(tokens[2*i])
		var err error = nil
		if i < len(tokens)/2-1 {
			value, err = strconv.Atoi(utils.ReadUntilLast(tokens[2*i+1]))
		} else {
			value, err = strconv.Atoi(tokens[2*i+1])
		}
		if err != nil {
			panic(tokens)
		}
		switch key {
		case "children":
			a.children = value
		case "cats":
			a.cats = value
		case "samoyeds":
			a.samoyeds = value
		case "pomeranians":
			a.pomeranians = value
		case "akitas":
			a.akitas = value
		case "vizslas":
			a.vizslas = value
		case "goldfish":
			a.goldfish = value
		case "trees":
			a.trees = value
		case "cars":
			a.cars = value
		case "perfumes":
			a.perfumes = value
		default:
			panic(tokens)
		}
	}
}

func linearSearch(candidatesPtr *[]Aunt, expectedAuntPtr *Aunt, exactRange bool) *Aunt {
	candidates := *candidatesPtr
	expectedAunt := *expectedAuntPtr

	worker := func(x int, y int) bool {
		if y == invalidValue {
			panic(y)
		}
		if x == invalidValue {
			return true
		}
		return x == y
	}
	workerGreater := func(x int, y int) bool {
		if y == invalidValue {
			panic(y)
		}
		if x == invalidValue {
			return true
		}
		return x > y
	}
	workerSmaller := func(x int, y int) bool {
		if y == invalidValue {
			panic(y)
		}
		if x == invalidValue {
			return true
		}
		return x < y
	}

	var matchingCandidates []Aunt
	for _, candidate := range candidates {
		matches := true

		// TODO: use reflection
		matches = matches && worker(candidate.children, expectedAunt.children)

		if exactRange {
			matches = matches && worker(candidate.cats, expectedAunt.cats)
		} else {
			matches = matches && workerGreater(candidate.cats, expectedAunt.cats)
		}
		matches = matches && worker(candidate.samoyeds, expectedAunt.samoyeds)

		if exactRange {
			matches = matches && worker(candidate.pomeranians, expectedAunt.pomeranians)
		} else {
			matches = matches && workerSmaller(candidate.pomeranians, expectedAunt.pomeranians)
		}

		matches = matches && worker(candidate.akitas, expectedAunt.akitas)
		matches = matches && worker(candidate.vizslas, expectedAunt.vizslas)

		if exactRange {
			matches = matches && worker(candidate.goldfish, expectedAunt.goldfish)
		} else {
			matches = matches && workerSmaller(candidate.goldfish, expectedAunt.goldfish)
		}

		if exactRange {
			matches = matches && worker(candidate.trees, expectedAunt.trees)
		} else {
			matches = matches && workerGreater(candidate.trees, expectedAunt.trees)
		}

		matches = matches && worker(candidate.cars, expectedAunt.cars)
		matches = matches && worker(candidate.perfumes, expectedAunt.perfumes)

		if matches {
			matchingCandidates = append(matchingCandidates, candidate)
		}
	}

	if len(matchingCandidates) == 0 {
		return nil
	} else if len(matchingCandidates) != 1 {
		panic(matchingCandidates)
	} else {
		return &matchingCandidates[0]
	}
}

func main() {
	lines := utils.GetLines("input")
	var candidates []Aunt

	for _, line := range lines {
		var aunt Aunt
		aunt.Parse(line)

		candidates = append(candidates, aunt)
	}

	var expectedAunt Aunt
	expectedAunt.children = 3
	expectedAunt.cats = 7
	expectedAunt.samoyeds = 2
	expectedAunt.pomeranians = 3
	expectedAunt.akitas = 0
	expectedAunt.vizslas = 0
	expectedAunt.goldfish = 5
	expectedAunt.trees = 3
	expectedAunt.cars = 2
	expectedAunt.perfumes = 1

	exactMatch := linearSearch(&candidates, &expectedAunt, true)
	fmt.Println(exactMatch.id)
	utils.WriteIntegerOutput(exactMatch.id, "1")

	rangeMatch := linearSearch(&candidates, &expectedAunt, false)
	fmt.Println(rangeMatch.id)
	utils.WriteIntegerOutput(rangeMatch.id, "2")
}
