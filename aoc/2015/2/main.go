package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

func getSurface(l int, w int, h int) int {
	return 2 * (l*w + w*h + h*l)
}

func getVolume(l int, w int, h int) int {
	return l * w * h
}

func getPerimeter(side1 int, side2 int) int {
	return 2.0 * (side1 + side2)
}

func main() {
	lines := utils.GetLines("input")

	totalSurface := 0
	totalRibbon := 0
	for _, line := range lines {
		tokens := strings.Split(line, "x")

		lStr, wStr, hStr := tokens[0], tokens[1], tokens[2]
		l, _ := strconv.ParseInt(lStr, 10, 64)
		w, _ := strconv.ParseInt(wStr, 10, 64)
		h, _ := strconv.ParseInt(hStr, 10, 64)
		totalSurface += getSurface(int(l), int(w), int(h))

		dimensions := []int{int(l), int(w), int(h)}
		sort.Ints(dimensions)

		totalSurface += dimensions[0] * dimensions[1]

		totalRibbon += getVolume(int(l), int(w), int(h))
		totalRibbon += getPerimeter(dimensions[0], dimensions[1])
	}
	fmt.Println("totalSurface = ", totalSurface)
	fmt.Println("totalRibbon = ", totalRibbon)

	utils.WriteLines([]string{fmt.Sprintf("%d", totalSurface)}, "output_1")
	utils.WriteLines([]string{fmt.Sprintf("%d", totalRibbon)}, "output_1")
}
