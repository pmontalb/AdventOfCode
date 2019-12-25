package main

import "fmt"

import "utils"

func test() {

}

func main() {
	p1 := FindNextPassword("hepxcrrq")
	fmt.Println(p1)
	utils.WriteSingleStringOutput(p1, "1")

	p2 := FindNextPassword(p1)
	fmt.Println(p2)
	utils.WriteSingleStringOutput(p2, "2")
}
