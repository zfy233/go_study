package main

import "fmt"

type interfaceTest interface {
	Add(x, y int) int
	Show(s string)
}

type Implement struct {
}

func (r *Implement) Add(x, y int) int {
	return x + y
}

func (r *Implement) Show(s string) {
	fmt.Println(s)
}

func main() {
	var a interfaceTest = &Implement{}
	z := a.Add(1, 2)
	fmt.Println(z)
}
