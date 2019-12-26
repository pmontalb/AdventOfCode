package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

func parseRelationships(relationships map[string]map[string]int, lines []string) []string {
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 11 {
			panic(tokens)
		}

		sign := 1
		if tokens[2] == "lose" {
			sign = -1
		}

		value, err := strconv.Atoi(tokens[3])
		if err != nil {
			panic(tokens)
		}

		if _, found := relationships[tokens[0]]; !found {
			relationships[tokens[0]] = make(map[string]int)
		}

		lastPerson := tokens[len(tokens)-1]
		lastPerson = lastPerson[:len(lastPerson)-1]
		relationships[tokens[0]][lastPerson] = sign * value
	}

	people := make([]string, len(relationships))
	i := 0
	for key := range relationships {
		people[i] = key
		i++
	}
	return people
}

func evaluatePermutation(permutation []string, relationships map[string]map[string]int) int {
	score := 0

	nPeople := len(permutation)
	//fmt.Printf("%d %s <-> %s: %d\n", 0, permutation[0], permutation[nPeople-1], relationships[permutation[0]][permutation[nPeople-1]])
	score += relationships[permutation[0]][permutation[nPeople-1]]

	for i := 0; i < nPeople-1; i++ {
		//fmt.Printf("%d %s <-> %s: %d\n", i, permutation[i], permutation[i+1], relationships[permutation[i]][permutation[i+1]])
		score += relationships[permutation[i]][permutation[i+1]]
		//fmt.Printf("%d %s <-> %s: %d\n", i, permutation[i+1], permutation[i], relationships[permutation[i+1]][permutation[i]])
		score += relationships[permutation[i+1]][permutation[i]]
	}

	//fmt.Printf("%d %s <-> %s: %d\n", nPeople-1, permutation[nPeople-1], permutation[0], relationships[permutation[nPeople-1]][permutation[0]])
	score += relationships[permutation[nPeople-1]][permutation[0]]

	return score
}

// TODO: use only circular permutations
//https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func findAllPermutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func findOptimalSeating(people []string, relationships map[string]map[string]int) (int, []string) {
	optimumScore := -math.MaxInt64
	var optimalSeating []string

	permutations := findAllPermutations(people)
	for _, permutation := range permutations {
		currentScore := evaluatePermutation(permutation, relationships)
		if currentScore > optimumScore {
			//fmt.Println(permutation, "*NEW optimum(", currentScore, ") <- max =", optimumScore)
			optimumScore = currentScore
			optimalSeating = permutation
		} else {
			//fmt.Println(permutation, "current target(", currentScore, ") | max =", optimumScore)
		}
	}

	return optimumScore, optimalSeating
}

func main() {
	lines := utils.GetLines("input")
	relationships := make(map[string]map[string]int)
	people := parseRelationships(relationships, lines)

	score, seating := findOptimalSeating(people, relationships)
	fmt.Println(score, seating)
	utils.WriteIntegerOutput(score, "1")

	me := "Paolo"
	relationships[me] = make(map[string]int)
	for _, person := range people {
		relationships[me][person] = 0
		relationships[person][me] = 0
	}
	people = append(people, me)

	score, seating = findOptimalSeating(people, relationships)
	fmt.Println(score, seating)
	utils.WriteIntegerOutput(score, "2")
}
