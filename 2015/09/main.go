package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"../utils"
)

// Path stores the graph info
type Path struct {
	start    string // only for reference
	end      string
	distance int // we will be sorting by this, having a map key'd by start city, and the values being this struct
}

// Paths is a collection of Path
type Paths []Path

func (paths Paths) Len() int           { return len(paths) }
func (paths Paths) Swap(i, j int)      { paths[i], paths[j] = paths[j], paths[i] }
func (paths Paths) Less(i, j int) bool { return paths[i].distance < paths[j].distance }

func getMinMaxPath(cities []string, distances map[string][]Path, isMin bool) ([]Path, int) {

	startIdx := 0
	idxIncrement := 1
	if !isMin {
		startIdx = len(cities) - 2 // the city itself must not be included!
		idxIncrement = -1
	}

	pathPerCity := make(map[string][]Path)

	targetPathLength := math.MaxInt64
	targetCity := ""
	if !isMin {
		targetPathLength = 0
	}
	for _, city := range cities {
		visitedCities := make(map[string]bool)

		lastVisitedCity := city
		visitedCities[lastVisitedCity] = true

		totalPathLength := 0
		for len(visitedCities) < len(cities) {
			// take the path to the shortest city
			pathIndex := startIdx
			for pathIndex < len(cities) {
				tryCity := distances[lastVisitedCity][pathIndex].end
				if _, found := visitedCities[tryCity]; !found {
					break
				}
				pathIndex += idxIncrement
			}
			if pathIndex >= len(cities) {
				fmt.Println(visitedCities)
				fmt.Println(cities)
				panic(pathIndex)
			}

			pathPerCity[city] = append(pathPerCity[city], distances[lastVisitedCity][pathIndex])
			totalPathLength += distances[lastVisitedCity][pathIndex].distance

			//fmt.Println("[", city, "]", lastVisitedCity, "->", distances[lastVisitedCity][pathIndex].end, "(", distances[lastVisitedCity][pathIndex].distance, "->", totalPathLength, ")")
			lastVisitedCity = distances[lastVisitedCity][pathIndex].end
			visitedCities[lastVisitedCity] = true
		}

		if isMin && totalPathLength < targetPathLength {
			targetPathLength = totalPathLength
			targetCity = city
		}
		if !isMin && totalPathLength > targetPathLength {
			targetPathLength = totalPathLength
			targetCity = city
		}
	}

	return pathPerCity[targetCity], targetPathLength
}

func main() {
	lines := utils.GetLines("input")

	distances := make(map[string][]Path)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 5 {
			panic(tokens)
		}
		distance, _ := strconv.Atoi(tokens[4])
		distances[tokens[0]] = append(distances[tokens[0]], Path{start: tokens[0], end: tokens[2], distance: distance})
		distances[tokens[2]] = append(distances[tokens[2]], Path{start: tokens[2], end: tokens[0], distance: distance})
	}

	var cities []string
	// sort by distance
	for startCity, paths := range distances {
		sort.Sort(Paths(paths))
		distances[startCity] = paths
		cities = append(cities, startCity)
	}

	_, minLength := getMinMaxPath(cities, distances, true)
	//fmt.Println(minPath)
	fmt.Println(minLength)
	utils.WriteIntegerOutput(minLength, "1")

	_, maxLength := getMinMaxPath(cities, distances, false)
	//fmt.Println(maxPath)
	fmt.Println(maxLength)
	utils.WriteIntegerOutput(maxLength, "2")
}
