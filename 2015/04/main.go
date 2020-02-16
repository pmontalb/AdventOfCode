package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"../utils"
)

const maxInteger int = 2<<32 - 1

func main() {
	lines := utils.GetLines("input")
	if len(lines) != 1 {
		panic(lines)
	}
	secretKey := lines[0]

	fiveZerosIdx := -1
	sixZerosIdx := -1
	for i := 0; i < maxInteger; i++ {
		key := secretKey + strconv.FormatInt(int64(i), 10)
		hasher := md5.New()
		hasher.Write([]byte(key))

		x := hex.EncodeToString(hasher.Sum(nil))
		if fiveZerosIdx == -1 && strings.HasPrefix(x, "00000") {
			fiveZerosIdx = i
			fmt.Println("Hash", x[:10], "... found at iteration", i)
		}
		if sixZerosIdx == -1 && strings.HasPrefix(x, "000000") {
			sixZerosIdx = i
			fmt.Println("Hash", x[:10], "... found at iteration", i)
		}
		if fiveZerosIdx != -1 && sixZerosIdx != -1 {
			break
		}
	}
	utils.WriteLines([]string{fmt.Sprintf("%d", fiveZerosIdx)}, "output_1")
	utils.WriteLines([]string{fmt.Sprintf("%d", sixZerosIdx)}, "output_2")
}
