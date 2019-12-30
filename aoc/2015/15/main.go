package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"../../../utils"
)

type Recipee struct {
	id         string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int

	teaspoons       int
	optimalTeaspons int
}

func (r *Recipee) Parse(line string) {
	tokens := strings.Split(line, " ")
	if len(tokens) != 11 {
		panic(tokens)
	}

	readUntilLastChar := func(token string) string {
		return token[:len(token)-1]
	}

	r.id = readUntilLastChar(tokens[0])
	r.capacity, _ = strconv.Atoi(readUntilLastChar(tokens[2]))
	r.durability, _ = strconv.Atoi(readUntilLastChar(tokens[4]))
	r.flavor, _ = strconv.Atoi(readUntilLastChar(tokens[6]))
	r.texture, _ = strconv.Atoi(readUntilLastChar(tokens[8]))
	r.calories, _ = strconv.Atoi(tokens[10])
}
func (r *Recipee) Floor() {
	if r.capacity < 0 {
		r.capacity = 0
	}
	if r.durability < 0 {
		r.durability = 0
	}
	if r.flavor < 0 {
		r.flavor = 0
	}
	if r.texture < 0 {
		r.texture = 0
	}
	if r.calories < 0 {
		r.calories = 0
	}
}

func mixRecipees(recipees *[]Recipee) Recipee {
	var mix Recipee

	for _, r := range *recipees {
		mix.capacity += r.capacity * r.teaspoons
		mix.durability += r.durability * r.teaspoons
		mix.flavor += r.flavor * r.teaspoons
		mix.texture += r.texture * r.teaspoons
		mix.calories += r.calories * r.teaspoons
	}
	mix.Floor()

	return mix
}

func evaluate(mix Recipee) int {
	score := 1
	score *= mix.capacity
	score *= mix.durability
	score *= mix.flavor
	score *= mix.texture

	return score
}

func optimize(recipees *[]Recipee, fixedCalories bool) int {
	optimum := -math.MaxInt64
	return optimizeWorker(0, recipees, optimum, fixedCalories)
}

// for now, simple grid search
func optimizeWorker(startIdx int, recipeesPtr *[]Recipee, currentOptimum int, fixedCalories bool) int {
	recipees := *recipeesPtr

	const maxTeaspoons = 100
	const fixedCaloriesAmount = 500
	nVariables := len(recipees)
	if startIdx >= nVariables-1 {
		totalTeaspoon := 0
		for i := 0; i < nVariables-1; i++ {
			totalTeaspoon += recipees[i].teaspoons
		}
		if totalTeaspoon > maxTeaspoons {
			return -math.MaxInt64
		}
		recipees[startIdx].teaspoons = maxTeaspoons - totalTeaspoon

		mix := mixRecipees(recipeesPtr)
		if fixedCalories && mix.calories != fixedCaloriesAmount {
			return -math.MaxInt64
		}
		return evaluate(mix)
	}

	currentLocalOptimum := 0
	for i := startIdx; i < nVariables-1; i++ {
		for t := 0; t < maxTeaspoons; t++ {
			recipees[i].teaspoons = t
			currentLocalOptimum = optimizeWorker(startIdx+1, recipeesPtr, currentOptimum, fixedCalories)

			totalTeaspoon := 0
			for _, r := range recipees {
				totalTeaspoon += r.teaspoons
			}
			if totalTeaspoon > maxTeaspoons {
				//fmt.Printf("%*s i(%d) t(%2d) totTsp(%d) <-- %v\n", 2*(startIdx+1), "*", i, t, totalTeaspoon, recipees)
				continue
			}

			if currentLocalOptimum > currentOptimum {
				for j := startIdx; j < nVariables; j++ {
					recipees[j].optimalTeaspons = recipees[j].teaspoons
				}
				// if startIdx == 0 {
				// 	fmt.Printf("%*s i(%d) t(%2d) *NEW obj(%12d) > %12d <-- %v\n", startIdx+1, "*", i, t, currentLocalOptimum, currentOptimum, recipees)
				// }

				currentOptimum = currentLocalOptimum
			} else {
				// if startIdx == 0 {
				// 	fmt.Printf("%*s i(%d) t(%d) obj(%d) <= %d <-- %v\n", startIdx+1, "*", i, t, currentLocalOptimum, currentOptimum, recipees)
				// }
			}
		}
	}

	for i := startIdx; i < nVariables; i++ {
		recipees[i].teaspoons = recipees[i].optimalTeaspons
	}
	return currentOptimum
}

func main() {
	lines := utils.GetLines("input")
	recipees := make([]Recipee, len(lines))
	for i, line := range lines {
		recipees[i].Parse(line)
	}

	optimize(&recipees, false)
	fmt.Println(evaluate(mixRecipees(&recipees)))
	utils.WriteIntegerOutput(evaluate(mixRecipees(&recipees)), "1")

	for i := 0; i < len(recipees); i++ {
		recipees[i].teaspoons = 0
	}
	optimize(&recipees, true)
	fmt.Println(evaluate(mixRecipees(&recipees)))
}
