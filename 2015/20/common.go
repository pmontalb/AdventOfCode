package main

import (
	"math"
	"sort"
)

func getDivisors(n int) []int {
	var divisors []int

	divisors = append(divisors, 1)

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if n/i != i {
				divisors = append(divisors, n/i)
			}
		}
	}

	if n > 1 {
		divisors = append(divisors, n)
	}
	sort.Ints(divisors)
	return divisors
}

func getDivisorsSum(n int) int {
	divisors := getDivisors(n)

	sum := 0
	for _, div := range divisors {
		sum += div
	}

	return sum
}

func countPresents(houseNumber int) int {
	sqrtN := int(math.Sqrt(float64(houseNumber)))

	nPresents := 1
	for i := 2; i <= sqrtN; i++ {
		if houseNumber%i == 0 {
			nPresents += i

			coFactor := houseNumber / i
			if coFactor != i {
				nPresents += coFactor
			}
		}
	}
	if houseNumber > 1 {
		nPresents += houseNumber
	}

	return 10 * nPresents
}

func countPresentsWithLimit(houseNumber int, limit int) int {
	/*
	 * The commented solution below is too slow, even with multiple threads.
	 * I've exploited the fact that this problem is simply to count:
	 * 1) the count of all divisors
	 * 2) the count of all divisors with quotient less than limit
	 * Therefore this function is a particular case of countPresents (which is just an alias of getDivisorsSum)
	 */

	// nPresents := 0
	// for i := 1; i <= houseNumber; i++ {
	// 	count := 0
	// 	for j := i; j <= houseNumber; j += i {
	// 		if count >= limit {
	// 			//fmt.Printf("Elf(%d) not visiting house(%d)(%d >= %d)\n", i, j, j-i, limit)
	// 			break
	// 		}
	// 		//fmt.Printf("Elf(%d) visiting house(%d)\n", i, j)
	// 		if j == houseNumber {
	// 			nPresents += i
	// 		}
	// 		count++
	// 	}
	// }
	// //fmt.Println()

	// return 11 * nPresent

	sqrtN := int(math.Sqrt(float64(houseNumber)))
	nPresents := 1
	if houseNumber > limit {
		nPresents = 0
	}
	for i := 2; i <= sqrtN; i++ {
		if houseNumber%i == 0 {
			modulus := houseNumber / i
			if modulus <= limit {
				nPresents += i
			}
			if modulus != i && i <= limit {
				nPresents += modulus
			}
		}
	}
	if houseNumber > 1 {
		nPresents += houseNumber
	}

	return 11 * nPresents
}
