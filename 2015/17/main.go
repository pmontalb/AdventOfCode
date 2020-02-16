package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"../utils"
)

func main() {
	lines := utils.GetLines("input")

	var containers []int
	for _, line := range lines {
		capacity, _ := strconv.Atoi(line)
		containers = append(containers, capacity)
	}

	sort.Ints(containers)

	combinations := *FindCombinations(&containers, 150)
	fmt.Println(len(combinations))
	utils.WriteIntegerOutput(len(combinations), "1")

	nMinContainers := math.MaxInt64
	for _, combination := range combinations {
		if len(combination) < nMinContainers {
			nMinContainers = len(combination)
		}
	}

	var combinationsWithMinContainers [][]int
	for _, combination := range combinations {
		if len(combination) == nMinContainers {
			combinationsWithMinContainers = append(combinationsWithMinContainers, combination)
		}
	}
	fmt.Println(len(combinationsWithMinContainers))
	utils.WriteIntegerOutput(len(combinationsWithMinContainers), "2")
}
