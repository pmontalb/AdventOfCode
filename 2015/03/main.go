package main

import (
	"fmt"
	"../utils"
)

// HouseLocation describe the location of the visited house
type HouseLocation struct {
	x int
	y int
}

func updatePosition(instruction string, loc *HouseLocation) {
	switch string(instruction) {
	case "^":
		loc.y++
	case "v":
		loc.y--
	case ">":
		loc.x++
	case "<":
		loc.x--
	}
}

func main() {
	lines := utils.GetLines("input")
	if len(lines) != 1 {
		panic(lines)
	}

	currentPosition := HouseLocation{0, 0}
	visitedHouses := make(map[HouseLocation]bool)
	visitedHouses[currentPosition] = true

	instructions := lines[0]
	for _, instruction := range instructions {
		updatePosition(string(instruction), &currentPosition)
		visitedHouses[currentPosition] = true
	}
	fmt.Println("Only Santa = ", len(visitedHouses))
	utils.WriteLines([]string{fmt.Sprintf("%d", len(visitedHouses))}, "output_1")

	if len(instructions)%2 != 0 {
		panic(len(instructions))
	}
	currentSantaPosition := HouseLocation{0, 0}
	currentRoboSantaPosition := HouseLocation{0, 0}
	visitedHouses = make(map[HouseLocation]bool)
	visitedHouses[currentSantaPosition] = true

	for i := 0; i < len(instructions)/2; i++ {
		updatePosition(string(instructions[2*i]), &currentSantaPosition)
		visitedHouses[currentSantaPosition] = true

		updatePosition(string(instructions[2*i+1]), &currentRoboSantaPosition)
		visitedHouses[currentRoboSantaPosition] = true
	}
	fmt.Println("With Robo Santa = ", len(visitedHouses))
	utils.WriteLines([]string{fmt.Sprintf("%d", len(visitedHouses))}, "output_2")
}
