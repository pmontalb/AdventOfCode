package main

import (
	"fmt"
	"sort"
	"strconv"

	"../utils"
)

func sum(weights []int) int {
	s := 0
	for _, w := range weights {
		s += w
	}
	return s
}

func main() {
	var weights []int
	lines := utils.GetLines("input")
	for _, line := range lines {
		w, err := strconv.Atoi(line)
		if err != nil {
			panic(0)
		}
		weights = append(weights, w)
	}

	sort.Ints(weights)

	nGroups := 3
	qe := OptimizeWorker(weights, sum(weights)/nGroups)
	fmt.Println(qe)
	utils.WriteIntegerOutput(qe, "1")

	nGroups = 4
	qe = OptimizeWorker(weights, sum(weights)/nGroups)
	fmt.Println(qe)
	utils.WriteIntegerOutput(qe, "2")
}
