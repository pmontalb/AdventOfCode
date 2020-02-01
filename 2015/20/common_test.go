package main

import (
	"math"
	"testing"
)

func TestAnalytical(t *testing.T) {
	for i := 1; i < 100; i++ {
		direct := countPresents(i)
		directWithLimit := countPresentsWithLimit(i, math.MaxInt64)
		analytical := 10 * getDivisorsSum(i)
		if direct != analytical {
			t.Errorf("i=%d, got %d, want %d (analytical)", i, direct, analytical)
		}
		if direct != directWithLimit/11*10 {
			t.Errorf("i=%d, got %d, want %d (analytical with no limit)", i, direct, directWithLimit)
		}
	}
}

func TestAnalyticalWithLimitOne(t *testing.T) { // if limit == 1, only one elf will deliver present
	for i := 2; i < 100; i++ {
		directWithOneLimit := countPresentsWithLimit(i, 1)
		if directWithOneLimit != 11*i {
			t.Errorf("i=%d, got %d, want %d (analytical with limit=1)", i, directWithOneLimit, 11*i)
		}
	}
}

func TestAnalyticalWithLimitTwo(t *testing.T) {
	expectedResults := []int{11, 33, 33, 66, 55, 99, 77, 132, 99, 165}
	for i := 1; i < 10; i++ {
		actual := countPresentsWithLimit(i, 2)
		expected := expectedResults[i-1]
		if actual != expected {
			t.Errorf("i=%d, got %d, want %d (analytical with limit=2)", i, actual, expected)
		}
	}
}

func TestAnalyticalWithLimitThree(t *testing.T) {
	expectedResults := []int{11, 33, 44, 66, 55, 121, 77, 132, 132, 165}
	for i := 1; i < 10; i++ {
		actual := countPresentsWithLimit(i, 3)
		expected := expectedResults[i-1]
		if actual != expected {
			t.Errorf("i=%d, got %d, want %d (AoC comparison)", i, actual, expected)
		}
	}
}

func TestAgainstAocResults(t *testing.T) {
	expectedResults := []int{10, 30, 40, 70, 60, 120, 80, 150, 130}
	for i := 0; i < len(expectedResults); i++ {
		actual := countPresents(i + 1)
		expected := expectedResults[i]
		if actual != expected {
			panic(i)
		}
	}
}
