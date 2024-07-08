package main

import (
	"errors"
	"fmt"
)


func Pembagian(nilai int,pembagi int ) (int, error){
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai/pembagi, nil
	}
}

func main(){
	fmt.Println(Pembagian(10, 0))
}