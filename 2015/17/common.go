package main

// assumption: containers is sorted
func findCombinationsWorker(offset int, containersPtr *[]int, currentStackPtr *[]int, target int) *[][]int {

	containers := *containersPtr
	currentStack := *currentStackPtr

	totalCapacity := 0
	for _, capacity := range currentStack {
		totalCapacity += capacity
	}

	var goodCombinations [][]int = nil
	if offset == len(containers) {
		if totalCapacity > target {
			panic(totalCapacity)
		}

		if totalCapacity == target {
			goodCombinations = append(goodCombinations, currentStack)
			return &goodCombinations
		}
		return nil
	}

	for i := offset; i < len(containers); i++ {
		newCapacity := totalCapacity + containers[i]
		if newCapacity >= target {
			if newCapacity == target {
				candidate := append(currentStack, containers[i])
				goodCombinations = append(goodCombinations, candidate)
			}

			continue
		}
		candidate := append(currentStack, containers[i])
		combinations := findCombinationsWorker(i+1, containersPtr, &candidate, target)
		if combinations != nil {
			for _, c := range *combinations {
				goodCombinations = append(goodCombinations, c)
			}
		}
	}

	return &goodCombinations
}
func FindCombinations(containers *[]int, target int) *[][]int {
	var currentStack []int = nil
	return findCombinationsWorker(0, containers, &currentStack, target)
}
