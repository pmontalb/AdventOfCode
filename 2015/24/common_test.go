package main

import (
	"sort"
	"strconv"
	"testing"

	"../utils"
)

func TestCase1(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	sort.Ints(weights)

	nGroups := 3
	if OptimizeWorker(weights, sum(weights)/nGroups) != 99 {
		t.Error()
	}
}

func TestCase2(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	sort.Ints(weights)

	nGroups := 4
	if OptimizeWorker(weights, sum(weights)/nGroups) != 44 {
		t.Error()
	}
}

func Benchmark(b *testing.B) {
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
	OptimizeWorker(weights, sum(weights)/nGroups)
}
