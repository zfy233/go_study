package main

import (
	"awesomeProject/entity"
	"fmt"
)

var a = entity.User{
	UserName: "zfy",
	Password: "123",
}

func main() {
	fmt.Println(a.UserName)
	fmt.Println(a.Password)
}
