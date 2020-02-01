package main

import (
	"runtime"
	"sort"
)

type OptimalConfiguration []int

func (configuration OptimalConfiguration) getQuantumEntanglement() int {
	quantumEntanglement := 1
	for _, element := range configuration {
		quantumEntanglement *= element
	}

	return quantumEntanglement
}

func (lhs OptimalConfiguration) less(rhs OptimalConfiguration) bool {
	qeI := lhs.getQuantumEntanglement()
	qeJ := rhs.getQuantumEntanglement()
	return qeI < qeJ
}

type OptimalConfigurations [][]int

func (qec OptimalConfigurations) Len() int      { return len(qec) }
func (qec OptimalConfigurations) Swap(i, j int) { qec[i], qec[j] = qec[j], qec[i] }
func (qec OptimalConfigurations) Less(i, j int) bool {
	lhs := OptimalConfiguration(qec[i])
	rhs := OptimalConfiguration(qec[j])
	return lhs.less(rhs)
}

/*
* After having tried a brute force approach for getting the exact configurations of all nGroups,
  discussions here [https://www.reddit.com/r/adventofcode/comments/3y1s7f/day_24_solutions/]
  suggested to focus only on the first group. In which case, it's just a matter of trying
  all possible combinations with a given sum.

  Before I was trying all possible sum, but: since you have to create a partition of the
  weights, and all subgroups have to sum the same, the only possibility is that this sum
  value is the average, which simplifies further the problem
*/
func OptimizeWorker(weights []int, targetSum int) int {

	for k := 1; true; k++ {
		combinationsWithFixedSum := generateCombinationWithFixedSum(weights, k, targetSum)
		if len(combinationsWithFixedSum) > 0 {
			// now sort partition by quantum entanglement and n packages of the first group
			optimalConfigurations := OptimalConfigurations(combinationsWithFixedSum)
			sort.Sort(optimalConfigurations)

			optimum := OptimalConfiguration(optimalConfigurations[0])
			qe := optimum.getQuantumEntanglement()

			return qe
		}
	}
	panic(0)
}

func nCk(n int, k int) int {
	ret := 1
	for i := 0; i < k; i++ {
		ret *= n - i
	}
	return ret
}
func generateCombinationWithFixedSum(pool []int, k int, targetSum int) [][]int {
	n := len(pool)
	if k > n {
		panic(0)
	}

	worker := func(retValue chan<- [][]int, start int, end int) {
		var localCombinations [][]int

		indices := make([]int, k)
		for i := 0; i < k; i++ {
			indices[i] = i
		}

		for counter := 0; counter < end; counter++ {
			i := k - 1
			done := true
			for ; i >= 0; i-- {
				if indices[i] != i+n-k {
					done = false
					break
				}
			}
			if done {
				break
			}
			indices[i]++
			for j := i + 1; j < k; j++ {
				indices[j] = indices[j-1] + 1
			}
			if counter < start {
				continue
			}

			localCombinations = append(localCombinations, []int{})
			localCombinations[len(localCombinations)-1] = make([]int, k)

			localSum := 0
			for j := 0; j < k; j++ {
				localCombinations[len(localCombinations)-1][j] = pool[indices[j]]
				localSum += pool[indices[j]]
			}
			if localSum != targetSum {
				localCombinations = localCombinations[:len(localCombinations)-1]
			}
		}
		retValue <- localCombinations
		return
	}

	nCPU := runtime.NumCPU()
	response := make(chan [][]int, nCPU)

	nCombinations := nCk(n, k) / (nCk(k, k))
	nOpPerThread := (nCombinations + nCPU) / nCPU

	start := 0
	end := start + nOpPerThread
	for n := 0; n < nCPU; n++ {
		go worker(response, start, end)
		start = end
		end += nOpPerThread
	}

	var combinations [][]int

	combinations = append(combinations, []int{})
	s := 0
	combinations[len(combinations)-1] = make([]int, k)
	for i := 0; i < k; i++ {
		combinations[len(combinations)-1][i] = pool[i]
		s += pool[i]
	}
	if s != targetSum {
		combinations = combinations[:len(combinations)-1]
	}

	for n := 0; n < nCPU; n++ {
		currentCombinations := <-response
		for _, c := range currentCombinations {
			combinations = append(combinations, c)
		}
	}
	return combinations
}
