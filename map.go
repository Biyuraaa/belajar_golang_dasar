package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"name":    "M Bimo Bayu Bagaskara",
		"nim":     "187221049",
		"jurusan": "Sistem Informasi",
	}

	fmt.Println(mahasiswa["name"])
	fmt.Println(mahasiswa["nim"])
	fmt.Println(mahasiswa["jurusan"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "M Bimo Bayu Bagaskara"
	book["wrong"] = "Salah"

	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)

}