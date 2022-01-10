package main

import (
	"fmt"
	"math/rand"
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
	for index, value := range words {
		fmt.Printf("%v %d| \n", value, index)
	}

	count := GetDataWordCount()
	list := make([]string, 0)
	for i := 0; i < count; i++ {
		rnd := rand.Intn(i + 1)
		indexData := GetRandomName(rnd)
		list = append(list, indexData)
	}

	last := list[len(list)-1]
	for index, item := range list {
		printed := "%d-%v -> "
		if last == item {
			printed = "%d-%v "
		}

		fmt.Printf(printed, index, item)
	}
}

func GetRandomName(index int) string {
	if index > GetDataWordCount() {
		panic("index out of range")
	}

	dataList := strings.Split(data, " ")
	var val = dataList[index]

	return val
}

func GetDataWordCount() int {
	return len(strings.Split(data, " "))
}

const data string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
