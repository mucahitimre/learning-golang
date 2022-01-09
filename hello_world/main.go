package main

import (
	"fmt"

	"example.com/m/v2/hello_world/mascot"
)

const (
	max_size = 5
	name     = "mico"
)

func main() {
	message := "hello world"
	println(message)
	fmt.Println(mascot.BestMascot())
	println(max_size)
	println(name)
}
