package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "hi content manager"
	fmt.Println(name)

	otherName := name
	if strings.EqualFold(name, otherName) {
		fmt.Println("name is equal otherName")
	}

	var containsName = name[:2]
	if strings.Contains(name, containsName) {
		fmt.Printf("'%v' is contains to '%v' \n", name, containsName)
	}

	var words = strings.Split(name, " ")
	for _, v := range words {
		fmt.Printf("%v | ", v)
	}
}
