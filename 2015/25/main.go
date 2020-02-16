package main

import (
	"fmt"
	"strconv"
	"strings"

	"../utils"
)

func main() {

	line := utils.GetFirstLine("input")
	tokens := strings.Split(line, " ")
	if len(tokens) != 19 {
		panic(tokens)
	}
	row, err1 := strconv.Atoi(tokens[16][:len(tokens[16])-1])
	if err1 != nil {
		panic(0)
	}
	col, err2 := strconv.Atoi(tokens[18][:len(tokens[18])-1])
	if err2 != nil {
		panic(0)
	}
	n := GetLinearizedIndex(row-1, col-1)
	code := 20151125
	for i := 0; i < n; i++ {
		code = GetNextCode(code)
	}
	fmt.Println(code)
	utils.WriteIntegerOutput(code, "1")
}
