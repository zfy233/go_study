package main

import "fmt"

func main() {
	c := make(chan int)
	go func(x, y int) {
		c <- x + y
		//fmt.Println(a)
	}(1, 2)
	fmt.Println(<-c)
}
