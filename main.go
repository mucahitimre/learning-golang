package main

import (
	"fmt"

	"example.com/m/v2/mascot"
)

func main() {
	message := "hello world"
	println(message)
	fmt.Println(mascot.BestMascot())
}
