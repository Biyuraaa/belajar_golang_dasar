package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	names := newMap("Bimo")
	fmt.Println(names["name"])

	dataKosong := newMap("")

	if dataKosong == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(dataKosong["name"])
	}
}