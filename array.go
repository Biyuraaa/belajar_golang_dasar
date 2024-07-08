package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "M"
	names[1] = "Bimo"
	names[2] = "Bayu"
	names[3] = "Bagaskara"

	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	var names2 = [3]string{
		"Atria",
		"Nur",
		"Farradina",
	}

	fmt.Println(len(names2))
	fmt.Println(names2)
}