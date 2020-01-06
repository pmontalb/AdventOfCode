package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"

	"../../../utils"
)

const maxGuess = 1000000 // not the most generic one!

func noLimitWorker(retValue chan<- int, start int, end int, target int) {
	i := start
	for ; i < end; i++ {
		if countPresents(i) > target {
			retValue <- i
			return
		}
	}

	retValue <- math.MaxInt64
}

func limitWorker(retValue chan<- int, start int, end int, target int, limit int) {
	i := start
	for ; i < end; i++ {
		if countPresentsWithLimit(i, limit) > target {
			retValue <- i
			return
		}
	}

	retValue <- math.MaxInt64
}

func dispatch(target int, worker func(chan<- int, int, int)) int {
	nCPU := runtime.NumCPU()
	response := make(chan int, nCPU)
	nOpPerThread := maxGuess / nCPU

	for n := 0; n < nCPU; n++ {
		start := n * nOpPerThread
		end := start + nOpPerThread
		go worker(response, start, end)
	}

	nMin := math.MaxInt64
	for n := 0; n < nCPU; n++ {
		currentN := <-response
		if currentN < nMin {
			nMin = currentN
		}
	}

	return nMin
}

func main() {

	const singleThreaded = false

	target, _ := strconv.Atoi(utils.GetFirstLine("input"))
	i := 0
	if singleThreaded {
		for ; i < maxGuess; i++ {
			if countPresents(i) > target {
				break
			}
		}
	} else {
		i = dispatch(target, func(retValue chan<- int, start int, end int) {
			noLimitWorker(retValue, start, end, target)
		})
	}
	fmt.Println(i)
	utils.WriteIntegerOutput(i, "1")

	j := i
	if singleThreaded {
		for ; j < math.MaxInt64; j++ {
			if countPresentsWithLimit(j, 50) > target {
				break
			}
		}
	} else {
		j = dispatch(target, func(retValue chan<- int, start int, end int) {
			limitWorker(retValue, start, end, target, 50)
		})
	}
	fmt.Println(j)
	utils.WriteIntegerOutput(j, "2")
}
