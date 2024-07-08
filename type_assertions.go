package main

import "fmt"

func random() interface{} {
	return 12
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case int :
		fmt.Println("Integer", value )
	case string :
		fmt.Println("String", value)
	default:
		fmt.Println("Unknown")
	}
}