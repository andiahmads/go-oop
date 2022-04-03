package main

import (
	"fmt"

	"github.com/andiahmads/go-oop/encapsulation"
)

func main() {

	user := encapsulation.NewUser("andiahmad", 25)
	fmt.Println(user.Age)
}
