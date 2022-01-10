package examplepackage

import "fmt"

var outSlice = []int{1, 2}

func ExampleSlice() {
	var slice []int
	PrintSlice("> Not definition slices", slice) // [] len:0 cap:0
	if slice == nil {
		fmt.Println("> Slice is null")
	}

	slice = make([]int, 8, 9)
	PrintSlice("> Defined slices", slice) // [0 0 0 0 0 0 0 0] len:8 cap:9

	slice = append(slice, 58, -25)
	PrintSlice("> Append slices", slice) // [0 0 0 0 0 0 0 0 58 -25] len: 10 cap: 18

	// struct: %v -> https://golangdocs.com/string-formatting-in-golang
	PrintSlice("> ", outSlice)   // [1 2] len: 2 cap: 2
	DetailedPrintSlice(outSlice) // [1 2] | [1 2] | []int{1, 2} len: 2 cap: 2
}

func PrintSlice(message string, x []int) {
	if message == "" || message == "> " {
		message = message + "Slice"
	}
	fmt.Printf("%s: %v len: %d cap: %d \n", message, x, len(x), cap(x))
}

func DetailedPrintSlice(x []int) {
	fmt.Printf("> Slice: %v | %+v | %#v len: %d cap: %d \n", x, x, x, len(x), cap(x))
}
