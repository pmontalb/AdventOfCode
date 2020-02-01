package main

import (
	"sort"
	"testing"
)

func TestFindCombinationWithFiveContainers(t *testing.T) {
	containers := []int{20, 15, 10, 5, 5}
	sort.Ints(containers)

	if len(*FindCombinations(&containers, 25)) != 4 {
		t.Errorf("got %d combinations, want %d", len(*FindCombinations(&containers, 25)), 4)
	}
}
