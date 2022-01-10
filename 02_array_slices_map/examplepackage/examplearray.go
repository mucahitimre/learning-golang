package examplepackage

import "fmt"

var outArray = [2]int{200, 404}

func ExampleArray() {
	var array = [3]string{"go", "java", "javascript"}
	notVarKeyword := [2]bool{true, false}

	// herhangi bir tanım verilmez ise üç tane default float32'li bir dizi oluşturur.
	notVarKeywordTheMake := make([]float32, 3)

	// %d:int | %s:string | %t:bool | %f:float32 -> https://golangdocs.com/string-formatting-in-golang
	fmt.Printf("> Arrays: \n%d \n%s \n%t \n%f \n", outArray, array, notVarKeyword, notVarKeywordTheMake)
	// [200 404] [go java javascript] [true false] [0.000000 0.000000 0.000000]
}
