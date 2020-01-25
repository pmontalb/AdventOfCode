package main

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"strconv"

	"../../../utils"
)

const nGroups = 3

func sum(weights *[]int) int {
	ret := 0
	for _, weight := range *weights {
		ret += weight
	}
	return ret
}
func findWeightsWithGivenSum(weights *[]int, target int, currentSolution *[][]int, currentCandidates *[]int) {
	s := sum(currentCandidates)
	//fmt.Printf("(%v) gives(%d)\n", currentCandidates, s)

	// found it
	if s == target {
		//fmt.Printf("sum(%d) found(%v)[%v]\n", s, currentCandidates, currentSolution)

		toAppend := make([]int, len(*currentCandidates))
		for i, c := range *currentCandidates {
			toAppend[i] = c
		}
		*currentSolution = append(*currentSolution, toAppend)
		return
	}
	if s > target {
		//fmt.Printf("sum(%d) >= tgt(%d) (%v)\n", s, target, currentCandidates)
		return // stop recursion
	}

	newCandidates := make([]int, len(*currentCandidates)+1)
	for i, c := range *currentCandidates {
		newCandidates[i] = c
	}
	for i := 0; i < len(*weights); i++ {
		remainingWeights := (*weights)[i+1:]

		newCandidates[len(*currentCandidates)] = (*weights)[i]
		findWeightsWithGivenSum(&remainingWeights, target, currentSolution, &newCandidates)
	}
}

type Partition [][]int

func (partition Partition) Len() int           { return len(partition) }
func (partition Partition) Swap(i, j int)      { partition[i], partition[j] = partition[j], partition[i] }
func (partition Partition) Less(i, j int) bool { return len(partition[i]) < len(partition[j]) }

func findPartitions(elements *[]int, tiling *[][]int, nGroups int, currentSolution *[][][]int, currentCandidates *[][]int) {
	if len(*currentCandidates) == nGroups {
		// first check that they can hold all elements
		totalLength := 0
		for _, tuple := range *currentCandidates {
			totalLength += len(tuple)
		}
		if totalLength != len(*elements) {
			return
		}

		// check that it contains all elements
		foundElements := make([]int, len(*elements))
		for i, element := range *elements {
			foundElements[i] = 0
			for _, tuple := range *currentCandidates {
				for _, candidate := range tuple {
					if candidate == element {
						foundElements[i]++
					}
				}
			}
		}
		if len(foundElements) != len(*elements) {
			return
		}
		for _, foundElementCount := range foundElements {
			if foundElementCount != 1 {
				return
			}
		}

		// this is a partition
		// check if this solution has already been entered
		//sort.Sort(Partition(*currentCandidates))
		// for _, tuple := range *currentSolution {
		// 	if areEqual2D(&tuple, currentCandidates) {
		// 		return
		// 	}
		// }
		//fmt.Printf("*NEW solution(%v)\n", *currentCandidates)

		toAppend := make([][]int, len(*currentCandidates))
		for i, c := range *currentCandidates {
			toAppend[i] = c
		}
		*currentSolution = append(*currentSolution, toAppend)
		return
	}
	if len(*currentCandidates) > nGroups {
		return // stop recursing
	}

	newCandidates := make([][]int, len(*currentCandidates)+1)
	for i, c := range *currentCandidates {
		newCandidates[i] = c
	}
	for _, tile := range *tiling {
		newCandidates[len(*currentCandidates)] = tile
		findPartitions(elements, tiling, nGroups, currentSolution, &newCandidates)
	}
}

type OptimalConfiguration [][]int

func (configuration *OptimalConfiguration) getQuantumEntanglement() int {
	quantumEntanglement := 1
	for _, element := range (*configuration)[0] {
		quantumEntanglement *= element
	}

	return quantumEntanglement
}

func (lhs *OptimalConfiguration) less(rhs *OptimalConfiguration) bool {
	if len((*lhs)[0]) == len((*rhs)[0]) {
		qeI := lhs.getQuantumEntanglement()
		qeJ := rhs.getQuantumEntanglement()
		return qeI < qeJ
	}
	return len((*lhs)[0]) < len((*rhs)[0])
}

type OptimalConfigurations [][][]int

func (qec OptimalConfigurations) Len() int      { return len(qec) }
func (qec OptimalConfigurations) Swap(i, j int) { qec[i], qec[j] = qec[j], qec[i] }
func (qec OptimalConfigurations) Less(i, j int) bool {
	lhs := OptimalConfiguration(qec[i])
	rhs := OptimalConfiguration(qec[j])
	return lhs.less(&rhs)
}

func test() {
	weights := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	sort.Ints(weights)

	var targetTuples [][]int
	findWeightsWithGivenSum(&weights, 20, &targetTuples, &[]int{})
	if len(targetTuples) == 0 {
		//fmt.Printf("sum(%d) no solution\n", s)
		panic(0)
	}

	//fmt.Println(weights)
	sort.Sort(Partition(targetTuples))
	// for _, s := range targetTuples {
	// 	fmt.Println(s)
	// }

	var partitions [][][]int
	findPartitions(&weights, &targetTuples, nGroups, &partitions, &[][]int{})
	if len(partitions) == 0 {
		//fmt.Printf("sum(%d) no partitions\n", s)
		panic(0)
	}

	// now sort partition by quantum entanglement and n packages of the first group
	optimalConfigurations := OptimalConfigurations(partitions)
	sort.Sort(optimalConfigurations)
	// fmt.Printf("sum(%d) solution:\n", 20)
	// for _, configuration := range optimalConfigurations {
	// 	fmt.Println("\t", configuration)
	// }

	optimum := OptimalConfiguration(optimalConfigurations[0])
	if optimum.getQuantumEntanglement() != 99 {
		panic(0)
	}
}

func optimizeWorker(weights *[]int, targetSum int) int {

	var targetTuples [][]int
	findWeightsWithGivenSum(weights, targetSum, &targetTuples, &[]int{})
	if len(targetTuples) == 0 {
		fmt.Printf("sum(%d) no solution\n", targetSum)
		return math.MaxInt64
	}

	//fmt.Println(weights)
	sort.Sort(Partition(targetTuples))
	// for _, s := range targetTuples {
	// 	fmt.Println(s)
	// }

	var partitions [][][]int
	findPartitions(weights, &targetTuples, nGroups, &partitions, &[][]int{})
	if len(partitions) == 0 {
		//fmt.Printf("sum(%d) no partitions\n", targetSum)
		return math.MaxInt64
	}

	// now sort partition by quantum entanglement and n packages of the first group
	optimalConfigurations := OptimalConfigurations(partitions)
	sort.Sort(optimalConfigurations)

	optimum := OptimalConfiguration(optimalConfigurations[0])
	qe := optimum.getQuantumEntanglement()
	fmt.Printf("sum(%d) solution [qe=%d]:\n", targetSum, qe)
	for _, configuration := range optimalConfigurations {
		fmt.Println("\t", configuration)
	}

	return qe
}

func optimize(retValue chan<- int, weights []int, targetSumStart int, targetSumEnd int, targetSumIncrement int) {

	var candidates []int
	for targetSum := targetSumStart; targetSum < targetSumEnd; targetSum += targetSumIncrement {
		candidate := optimizeWorker(&weights, targetSum)
		candidates = append(candidates, candidate)
		fmt.Printf("sum(%d) candidate(%d)\n", targetSum, candidate)
	}

	min := math.MaxInt64
	for _, candidate := range candidates {
		if candidate < min {
			min = candidate
		}
	}
	retValue <- min
}

func dispatch(weights *[]int, worker func(chan<- int, []int, int, int, int)) int {
	maxSum := sum(weights)

	nCPU := runtime.NumCPU()
	response := make(chan int, nCPU)
	//nOpPerThread := maxSum / nCPU

	startSum := 1
	for n := 0; n < nCPU; n++ {
		start := startSum + n
		end := maxSum
		go worker(response, *weights, start, end, nCPU)
	}

	optimalQe := math.MaxInt64
	for n := 0; n < nCPU; n++ {
		currentQe := <-response
		if currentQe < optimalQe {
			optimalQe = currentQe
			return currentQe // don't wait for the others!
		}
	}

	return optimalQe
}

func main() {
	test()

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

	//dispatch(&weights, optimize)
	optimizeWorker(&weights, 150)
	// for i := 0; i < sum(&weights); i++ {
	// 	optimizeWorker(&weights, i)
	// }
}
