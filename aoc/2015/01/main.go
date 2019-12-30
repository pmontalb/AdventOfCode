package main

import (
	"fmt"
	"../../../utils"
)

func isCharacterValid(char string) bool {
	switch char {
	case
		"(",
		")":
		return true
	}
	return false
}

func main() {
	lines := utils.GetLines("input")
	if len(lines) != 1 {
		panic(len(lines))
	}

	instructions := lines[0]
	position := 0
	firstPosMinus1 := -1

	for i, instruction := range instructions {
		instr := string(instruction)
		if !isCharacterValid(instr) {
			panic(instr)
		}
		if instr == "(" {
			position++
			//fmt.Println("Going up to", position)
		} else {
			position--
			//fmt.Println("Going down to", position)
		}

		if position == -1 && firstPosMinus1 == -1 {
			firstPosMinus1 = i + 1
			fmt.Println("Reached", position, "after", firstPosMinus1, "iterations")
		}
	}
	fmt.Println("Position =", position)

	utils.WriteLines([]string{fmt.Sprintf("%d", position)}, "output_1")
	utils.WriteLines([]string{fmt.Sprintf("%d", firstPosMinus1)}, "output_2")
}
