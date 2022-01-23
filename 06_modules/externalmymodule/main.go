package externalmymodule

import "fmt"

func init() {
	fmt.Println("My external module")
}

func GetMyData() string {

	return "Foo"
}
