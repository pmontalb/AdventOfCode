package main

import (
	"fmt"
	"utils"
)

func isNiceString(str string) bool {
	vowelCount := 0
	hasDoubleLetter := false
	hasBlackListStrings := false
	for i := 0; i < len(str); i++ {
		token := string(str[i])
		switch token {
		case "a":
			fallthrough
		case "e":
			fallthrough
		case "i":
			fallthrough
		case "o":
			fallthrough
		case "u":
			vowelCount++
		}

		if i < len(str)-1 {
			hasDoubleLetter = hasDoubleLetter || str[i] == str[i+1]

			switch string(str[i : i+2]) {
			case "ab":
				fallthrough
			case "cd":
				fallthrough
			case "pq":
				fallthrough
			case "xy":
				hasBlackListStrings = true
			}
		}
	}

	return vowelCount >= 3 && hasDoubleLetter && !hasBlackListStrings
}

func isNicerString(str string) bool {
	hasNonOverlappingPair := false
	hasRepeatingPairWithOneLetterSeparation := false

	for i := 0; i < len(str)-2; i++ {
		pair := string(str[i : i+2])
		for j := i + 2; j < len(str)-1; j++ {
			nextPair := string(str[j : j+2])
			if pair == nextPair {
				hasNonOverlappingPair = true
			}
		}
	}
	if !hasNonOverlappingPair {
		return false
	}

	for i := 0; i < len(str)-2; i++ {
		token := string(str[i])
		nextToken := string(str[i+2])
		if token == nextToken {
			hasRepeatingPairWithOneLetterSeparation = true
		}
	}
	if !hasRepeatingPairWithOneLetterSeparation {
		return false
	}

	return true
}

func main() {
	niceStringCount := 0
	nicerStringCount := 0
	lines := utils.GetLines("input")
	for _, line := range lines {
		if isNiceString(line) {
			niceStringCount++
		}
		if isNicerString(line) {
			nicerStringCount++
		}
	}
	fmt.Println(niceStringCount)
	utils.WriteLines([]string{fmt.Sprintf("%d", niceStringCount)}, "output_1")

	fmt.Println(nicerStringCount)
	utils.WriteLines([]string{fmt.Sprintf("%d", nicerStringCount)}, "output_2")
}
