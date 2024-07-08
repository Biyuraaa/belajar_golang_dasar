package main

import "fmt"

func main() {
	names := [...]string{"M", "Bimo", "Bayu", "Bagaskara"}
	slice := names[0:3]

	fmt.Println(slice[0:2])

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	names2 := [...]string{"Atria", "Nur", "Farradina"}
	slice2 := names2[0:2]

	slice3 := append(slice2, "Farradina")

	fmt.Println(slice3)
}