package main

import "fmt"

type zzz string

func (z *zzz) test() string {
	fmt.Println("123")
	return "123"
}

func main() {
	var b zzz
	a := b.test()
	fmt.Println(a)
}
