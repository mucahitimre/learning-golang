package examplepackage

import (
	"fmt"
)

func ExampleMap() {
	var dictionary map[int]string
	if dictionary == nil {
		dictionary = make(map[int]string, 5)
	}
	dictionary[0] = "zero"
	fmt.Println(dictionary[0])

	langDictionary := make(map[string]string)
	langDictionary["Bangladesh"] = "Bengali"
	langDictionary["Germany"] = "German"
	langDictionary["India"] = "Telugu"
	langDictionary["France"] = "French"
	langDictionary["Korea"] = "Korean"

	keys := make([]string, 0, len(langDictionary))
	values := make([]string, 0, len(langDictionary))
	for k := range langDictionary {
		keys = append(keys, k)
		values = append(values, langDictionary[k])
	}

	fmt.Printf("values: %v\n", values) // values: [Bengali German Telugu French Korean]
	fmt.Printf("keys: %v\n", keys)     // keys: [Bangladesh Germany India France Korea]
}
