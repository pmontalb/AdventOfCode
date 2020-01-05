package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"../../../utils"
)

func getDistinctMolecules(molecule string, tokenReplacementMap map[string][]string) map[string]bool {
	differentMolecules := make(map[string]bool)

	for token, replacements := range tokenReplacementMap {
		moleculeCopy := molecule
		for i := 0; i < len(molecule)-len(token)+1; i++ {
			if molecule[i:i+len(token)] == token {
				for _, replacement := range replacements {
					// replace it with replacement
					newMolecule := moleculeCopy[:i]
					newMolecule += replacement
					newMolecule += moleculeCopy[i+len(token):]
					differentMolecules[newMolecule] = true

					//fmt.Printf("token(%s) found at idx(%d): replacing with(%s)\n", token, i, replacement)
				}
			}
		}
	}

	//fmt.Println(differentMolecules)
	return differentMolecules
}

// after trying a DFS for visualizing the actual path - taking too long - I've read there's only one possible solution.
//Therefore I can just sequentially replace all tokens until I get an electron
// Source: https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/
func transformMolecule(molecule string, tokenReplacementInverseMap map[string][]string, electronReplacements []string) int {
	count := 0

	transformedMolecule := molecule

	r := rand.New(rand.NewSource(1234))
	for true {
		hasReplaced := false
		for token, replacements := range tokenReplacementInverseMap {
			moleculeCopy := transformedMolecule
			for i := 0; i < len(transformedMolecule)-len(token)+1; i++ {
				if transformedMolecule[i:i+len(token)] == token {
					newMolecule := moleculeCopy[:i]
					newMolecule += replacements[r.Int()%len(replacements)]
					newMolecule += moleculeCopy[i+len(token):]
					transformedMolecule = newMolecule

					for _, et := range electronReplacements {
						if transformedMolecule == et {
							return count + 2
						}
					}
					count++
					hasReplaced = true
					break
				}
			}
		}

		if !hasReplaced {
			count = 0
			transformedMolecule = molecule
		}
	}
	return math.MaxInt64
}

func test() {
	tokenReplacementMap := make(map[string][]string)
	tokenReplacementMap["H"] = []string{"HO", "OH"}
	tokenReplacementMap["O"] = []string{"HH"}
	tokenReplacementMap["e"] = []string{"H", "O"}

	tokenReplacementInverseMap := make(map[string][]string)
	tokenReplacementInverseMap["HO"] = []string{"H"}
	tokenReplacementInverseMap["OH"] = []string{"H"}
	tokenReplacementInverseMap["HH"] = []string{"O"}

	electronReplacements := []string{"H", "O"}
	moleculeStr := "HOH"

	count := len(getDistinctMolecules(moleculeStr, tokenReplacementMap))
	if count != 4 {
		panic(count)
	}

	count = transformMolecule(moleculeStr, tokenReplacementInverseMap, electronReplacements)
	if count != 3 {
		panic(count)
	}

	moleculeStr = "HOHOHO"

	count = len(getDistinctMolecules(moleculeStr, tokenReplacementMap))
	if count != 7 {
		panic(count)
	}

	count = transformMolecule(moleculeStr, tokenReplacementInverseMap, electronReplacements)
	if count != 6 {
		panic(count)
	}
}

func main() {
	test()

	lines := utils.GetLines("input")
	tokenReplacementMap := make(map[string][]string)
	tokenReplacementInverseMap := make(map[string][]string)
	var electronReplacements []string
	lastIdx := 0
	for i, line := range lines {
		if line == "" {
			lastIdx = i
			break
		}

		tokens := strings.Split(line, "=>")
		if len(tokens) != 2 {
			panic(tokens)
		}
		tokenReplacementMap[strings.TrimSpace(tokens[0])] = append(tokenReplacementMap[strings.TrimSpace(tokens[0])], strings.TrimSpace(tokens[1]))

		if strings.TrimSpace(tokens[0]) == "e" {
			electronReplacements = append(electronReplacements, strings.TrimSpace(tokens[1]))
		} else {
			tokenReplacementInverseMap[strings.TrimSpace(tokens[1])] = append(tokenReplacementInverseMap[strings.TrimSpace(tokens[1])], strings.TrimSpace(tokens[0]))
		}
	}

	moleculeStr := lines[lastIdx+1]

	count := len(getDistinctMolecules(moleculeStr, tokenReplacementMap))
	fmt.Println(count)
	utils.WriteIntegerOutput(count, "1")

	count = transformMolecule(moleculeStr, tokenReplacementInverseMap, electronReplacements)
	fmt.Println(count)
	utils.WriteIntegerOutput(count, "2")
}
