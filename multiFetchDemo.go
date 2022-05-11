package main

import "fmt"

//var aa string = "123"
//
//const bb = 123
//
//var (
//	cc = 1
//	dd = 2
//)
//
//const (
//	ee = 1
//	ff = 2
//)
//
//var s string

type PointerTest struct {
	x int
	y string
}

func main() {
	p := PointerTest{
		x: 1,
		y: "zzz",
	}
	pointer := &p
	fmt.Println(&pointer.y)
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
