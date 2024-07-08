package main

import "fmt"

func main() {
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)


	benar := 2 == 1+1;

	fmt.Println("Benar = ", benar)

	salah := 2 == 1;

	fmt.Println("Salah = ", salah)
}