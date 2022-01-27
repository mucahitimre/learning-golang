package main

import (
	"Learngo/01_hello_world/mascot"
	"fmt"
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
