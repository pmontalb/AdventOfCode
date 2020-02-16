package main

import (
	"fmt"
	"strconv"
	"strings"
	"../utils"
)

// GridPoint describes points in the grid
type GridPoint struct {
	x int
	y int
}

// Instruction holds the instructio from input
type Instruction struct {
	instructionType string
	startPoint      GridPoint
	endPoint        GridPoint
}

// Parse parses the instruction from input
func (instr *Instruction) Parse(line string) {
	tokens := strings.Split(line, " ")
	if len(tokens) > 5 {
		panic(tokens)
	}

	firstCoordinateIdx := 1
	secondCoordinateIdx := 3
	if tokens[0] == "turn" {
		firstCoordinateIdx++
		secondCoordinateIdx++

		if tokens[1] == "on" {
			instr.instructionType = "on"
		} else if tokens[1] == "off" {
			instr.instructionType = "off"
		} else {
			panic(tokens)
		}
	} else if tokens[0] == "toggle" {
		instr.instructionType = "toggle"
	} else {
		panic(tokens)
	}

	startCoordinateStr := strings.Split(tokens[firstCoordinateIdx], ",")
	if len(startCoordinateStr) != 2 {
		panic(tokens)
	}
	instr.startPoint.x, _ = strconv.Atoi(startCoordinateStr[0])
	instr.startPoint.y, _ = strconv.Atoi(startCoordinateStr[1])

	endCoordinateStr := strings.Split(tokens[secondCoordinateIdx], ",")
	if len(endCoordinateStr) != 2 {
		panic(tokens)
	}
	instr.endPoint.x, _ = strconv.Atoi(endCoordinateStr[0])
	instr.endPoint.y, _ = strconv.Atoi(endCoordinateStr[1])
}

func process(instr Instruction, grid *[][]int) {

	worker := func(light *int) { /* noop */ }
	switch instr.instructionType {
	case "on":
		worker = func(light *int) { *light = 1 }
	case "off":
		worker = func(light *int) { *light = 0 }
	case "toggle":
		worker = func(light *int) { *light = 1 - *light }
	}

	for i := instr.startPoint.x; i <= instr.endPoint.x; i++ {
		for j := instr.startPoint.y; j <= instr.endPoint.y; j++ {
			worker(&((*grid)[i][j]))
		}
	}
}

func processBrightness(instr Instruction, grid *[][]int) {

	worker := func(light *int) { /* noop */ }
	switch instr.instructionType {
	case "on":
		worker = func(light *int) { *light++ }
	case "off":
		worker = func(light *int) {
			*light--
			if *light < 0 {
				*light = 0
			}
		}
	case "toggle":
		worker = func(light *int) { *light += 2 }
	}

	for i := instr.startPoint.x; i <= instr.endPoint.x; i++ {
		for j := instr.startPoint.y; j <= instr.endPoint.y; j++ {
			worker(&((*grid)[i][j]))
		}
	}
}

func main() {
	const dimX = 1000
	const dimY = 1000

	grid := make([][]int, dimX)
	for i := 0; i < dimX; i++ {
		grid[i] = make([]int, dimY)
	}

	lines := utils.GetLines("input")
	for _, line := range lines {
		instr := Instruction{}
		instr.Parse(line)

		process(instr, &grid)
	}

	onCount := 0
	for i := 0; i < dimX; i++ {
		for j := 0; j < dimY; j++ {
			onCount += grid[i][j]
		}
	}
	fmt.Println("onCount", onCount)
	utils.WriteLines([]string{strconv.Itoa(onCount)}, "output_1")

	// reset
	for i := 0; i < dimX; i++ {
		for j := 0; j < dimY; j++ {
			grid[i][j] = 0
		}
	}

	for _, line := range lines {
		instr := Instruction{}
		instr.Parse(line)

		processBrightness(instr, &grid)
	}

	totalBrightness := 0
	for i := 0; i < dimX; i++ {
		for j := 0; j < dimY; j++ {
			totalBrightness += grid[i][j]
		}
	}
	fmt.Println("totalBrightness", totalBrightness)
	utils.WriteLines([]string{strconv.Itoa(totalBrightness)}, "output_2")
}
