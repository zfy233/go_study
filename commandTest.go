package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//s := os.Args[0]
	//printCommand()
	//println(s)
	testPrintf()
}

func printCommand() {
	for i, value := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i) + " " + value)
	}
}

func testPrintf() {
	fmt.Printf("zzz%v", "aaa")
	fmt.Println()
	fmt.Println(123)
}
