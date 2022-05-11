package main

import "fmt"

func main() {
	//c1 := sha256.Sum256([]byte("x"))
	//c2 := sha256.Sum256([]byte("X"))
	//number := 0
	//for i, _ := range c1 {
	//	if c1[i] != c2[i] {
	//		number++
	//	}
	//}
	//fmt.Println(c1)
	//fmt.Println(c2)
	//fmt.Println(len(c1))
	//fmt.Println(number)
	//c := make(map[int]int)
	//c[0] = 1
	//testMap(c)
	//v, ok := c[1]
	//if ok {
	//	fmt.Println("outside OK")
	//	fmt.Println(v)
	//}
	//fmt.Printf("Type%T", c)
	//fmt.Printf("Type%T", pointer)

	c := make([]int, 0, 10)
	c = append(c, 1, 2)
	fmt.Println(&c[1])
	testSlice(c)
}

func testSlice(c []int) {
	fmt.Println(&c[1])
}

func testChan(c chan int) {
	fmt.Println(&c)
}

func testMap(c map[int]int) {
	_, ok := c[0]
	if ok {
		fmt.Println("OK")
	} else {
		fmt.Println("不包含c[0]")
	}
	c[1] = 1
}

func reverse(s []int) {
	for i := range s {
		if i < len(s)/2 {
			s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		}
	}
}

func delete(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = append(s[:i+1], s[i+2:]...)
		}
	}
	return s
}
