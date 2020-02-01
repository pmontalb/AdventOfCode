package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"../../../utils"
)

func digitiseWithPowerOfTen(numberString string) []int {
	number, _ := strconv.ParseFloat(numberString, 64)
	powerOfTen := math.Pow10(len(numberString) - 1)

	digits := make([]int, len(numberString))
	for i := 0; i < len(numberString); i++ {
		digits[i] = int(number / powerOfTen)

		number -= float64(digits[i]) * powerOfTen
		powerOfTen /= 10
	}

	return digits
}

func digitiseWithStrConv(numberString string) []string {

	digits := make([]string, len(numberString))
	for i := 0; i < len(numberString); i++ {
		digits[i] = string(numberString[i])
	}

	return digits
}

func lookAndSay(numberString string) string {

	var output bytes.Buffer
	digits := digitiseWithStrConv(numberString)
	for i := 0; i < len(digits); i++ {
		j := i + 1
		for ; j < len(digits); j++ {
			if digits[j] != digits[i] {
				//fmt.Println(i, j, digits[i], digits[j])
				break
			}
		}
		currentCount := j - i

		output.WriteString(strconv.Itoa(currentCount))
		output.WriteByte(numberString[i])
		//fmt.Println(i, j, "digit", string(numberString[i]), "[", digits[i], "] found", currentCount, "times ->", strconv.Itoa(currentCount)+string(numberString[i]))

		i = j - 1
	}

	return output.String()
}

func test() {
	if lookAndSay("1") != "11" {
		panic("1")
	}
	if lookAndSay("11") != "21" {
		panic("11")
	}
	if lookAndSay("21") != "1211" {
		panic("21")
	}
	if lookAndSay("1211") != "111221" {
		panic("1211")
	}
	if lookAndSay("111221") != "312211" {
		panic("111221")
	}

	x := "1"
	for i := 0; i < 5; i++ {
		x = lookAndSay(x)
	}
	if x != "312211" {
		panic("loop")
	}

	t := lookAndSay("1113122113121113222113")
	if t != "31131122211311123113322113" {
		panic("test")
	}
}

func main() {
	test()

	start := utils.GetFirstLine("input")
	for i := 0; i < 40; i++ {
		start = lookAndSay(start)
	}
	fmt.Println(len(start))
	utils.WriteIntegerOutput(len(start), "1")

	for i := 0; i < 10; i++ { // further 10 times
		start = lookAndSay(start)
	}
	fmt.Println(len(start))
	utils.WriteIntegerOutput(len(start), "2")
}
