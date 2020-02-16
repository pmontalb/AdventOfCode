package main

import (
	"fmt"
	"time"

	"../utils"
)

type Light struct {
	state int
}

func (l *Light) Parse(r rune) {
	switch r {
	case '#':
		l.state = 1
	case '.':
		l.state = 0
	default:
		panic(r)
	}
}
func (l *Light) Print() {
	if l.state == 1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}

type LightGrid [][]Light

func (l LightGrid) Count() int {
	count := 0
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l[i]); j++ {
			count += l[i][j].state
		}
	}

	return count
}
func (l LightGrid) Print(maxRow int) {
	for i := 0; i < maxRow; i++ {
		for j := 0; j < len(l[i]); j++ {
			l[i][j].Print()
		}
		fmt.Printf("\n")
	}
}
func (l LightGrid) Evolve(steps int) {
	lCopy := make([][]Light, len(l))

	for i := 0; i < len(l); i++ {
		lCopy[i] = make([]Light, len(l[i]))
		for j := 0; j < len(l[i]); j++ {
			lCopy[i][j] = l[i][j]
		}
	}
	lCopyGrid := LightGrid(lCopy)

	new := &l
	old := &lCopyGrid
	for i := 0; i < steps; i++ {
		evolve(new, old)
		old, new = new, old
	}

	if steps%2 == 0 { // we need to copy old into new
		copy(l, lCopy)
	}
}

func evolve(lightsPtrNew *LightGrid, lightsPtrOld *LightGrid) {
	lights := *lightsPtrNew
	lightsOld := *lightsPtrOld

	evolveWorker := func(lights *Light, lightsOld *Light, nNeighborhoodsOn int) {
		if lightsOld.state == 1 {
			if nNeighborhoodsOn != 2 && nNeighborhoodsOn != 3 {
				lights.state = 0
			} else {
				lights.state = 1
			}
		} else {
			if nNeighborhoodsOn == 3 {
				lights.state = 1
			} else {
				lights.state = 0
			}
		}
	}

	// process top row
	i := 0
	for j := 0; j < len(lights[0]); j++ {
		if j == 0 {
			// process top left corner
			nNeighborhoodsOn :=
				lightsOld[i][j+1].state +
					lightsOld[i+1][j].state + lightsOld[i+1][j+1].state

			evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
		} else if j == len(lights[0])-1 {
			// process top right corner
			nNeighborhoodsOn :=
				lightsOld[i][j-1].state +
					lightsOld[i+1][j-1].state + lightsOld[i+1][j].state

			evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
		} else {
			nNeighborhoodsOn :=
				lightsOld[i][j-1].state + lightsOld[i][j+1].state +
					lightsOld[i+1][j-1].state + lightsOld[i+1][j].state + lightsOld[i+1][j+1].state

			evolveWorker(&lights[0][j], &lightsOld[0][j], nNeighborhoodsOn)
		}
	}

	// process bottom row
	i = len(lights) - 1
	for j := 0; j < len(lights[0]); j++ {
		if j == 0 {
			// process bottom left corner
			nNeighborhoodsOn :=
				lightsOld[i-1][j].state + lightsOld[i-1][j+1].state +
					lightsOld[i][j+1].state

			evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
		} else if j == len(lights[0])-1 {
			// process bottom right corner
			nNeighborhoodsOn :=
				lightsOld[i-1][j-1].state + lightsOld[i-1][j].state +
					lightsOld[i][j-1].state

			evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
		} else {
			nNeighborhoodsOn :=
				lightsOld[i-1][j-1].state + lightsOld[i-1][j].state + lightsOld[i-1][j+1].state +
					lightsOld[i][j-1].state + lightsOld[i][j+1].state

			evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
		}
	}

	// process center
	for i = 1; i < len(lights)-1; i++ {
		for j := 0; j < len(lights[i]); j++ {
			if j == 0 {
				nNeighborhoodsOn :=
					lightsOld[i-1][j].state + lightsOld[i-1][j+1].state +
						lightsOld[i][j+1].state +
						lightsOld[i+1][j].state + lightsOld[i+1][j+1].state

				evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
			} else if j == len(lights[i])-1 {
				nNeighborhoodsOn :=
					lightsOld[i-1][j-1].state + lightsOld[i-1][j].state +
						lightsOld[i][j-1].state +
						lightsOld[i+1][j-1].state + lightsOld[i+1][j].state

				evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
			} else {
				// count neighborhoods
				nNeighborhoodsOn :=
					lightsOld[i-1][j-1].state + lightsOld[i-1][j].state + lightsOld[i-1][j+1].state +
						lightsOld[i][j-1].state + lightsOld[i][j+1].state +
						lightsOld[i+1][j-1].state + lightsOld[i+1][j].state + lightsOld[i+1][j+1].state

				evolveWorker(&lights[i][j], &lightsOld[i][j], nNeighborhoodsOn)
			}
		}
	}
}

func test() {
	var lines []string
	lines = append(lines, ".#.#.#")
	lines = append(lines, "...##.")
	lines = append(lines, "#....#")
	lines = append(lines, "..#...")
	lines = append(lines, "#.#..#")
	lines = append(lines, "####..")

	lights := make([][]Light, len(lines))

	for i, line := range lines {
		lights[i] = make([]Light, len(line))
		for j, r := range line {
			lights[i][j].Parse(r)
		}
	}

	grid := LightGrid(lights)

	grid.Evolve(1)
	grid.Evolve(1)
	grid.Evolve(1)
	grid.Evolve(1)
	if grid.Count() != 4 {
		panic(1)
	}

	// reset
	for i, line := range lines {
		for j, r := range line {
			lights[i][j].Parse(r)
		}
	}
	grid.Evolve(4)
	if grid.Count() != 4 {
		panic(1)
	}
}

func main() {
	test()

	lines := utils.GetLines("input")

	lights := make([][]Light, len(lines))

	for i, line := range lines {
		lights[i] = make([]Light, len(line))
		for j, r := range line {
			lights[i][j].Parse(r)
		}
	}

	grid := LightGrid(lights)

	const steps = 100
	const animateFirst = false
	if animateFirst {
		const waitMs = 25
		rowToPrint := len(lights) / 2

		for i := 0; i < steps; i++ {
			time.Sleep(waitMs * time.Millisecond)

			fmt.Printf("\033[0;0H")

			grid.Evolve(1)
			grid.Print(rowToPrint)
		}
	} else {
		grid.Evolve(steps)
	}

	fmt.Println(grid.Count())
	utils.WriteIntegerOutput(grid.Count(), "1")

	// now re-run it but with 4 corners always on
	for i, line := range lines {
		for j, r := range line {
			lights[i][j].Parse(r)
		}
	}

	setCornersOn := func() {
		lights[0][0].state = 1
		lights[0][len(lights)-1].state = 1
		lights[len(lights)-1][0].state = 1
		lights[len(lights)-1][len(lights)-1].state = 1
	}

	setCornersOn()

	const animateSecond = false
	if animateSecond {
		const waitMs = 25
		rowToPrint := len(lights) / 2

		for i := 0; i < steps; i++ {
			time.Sleep(waitMs * time.Millisecond)

			fmt.Printf("\033[0;0H")

			grid.Evolve(1)
			setCornersOn()

			grid.Print(rowToPrint)
		}
	} else {
		for i := 0; i < steps; i++ {
			grid.Evolve(1)
			setCornersOn()
		}
	}
	fmt.Println(grid.Count())
	utils.WriteIntegerOutput(grid.Count(), "2")
}
