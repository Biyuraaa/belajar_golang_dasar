package main

import (
	"belajar_golang_dasar/database"
	_ "belajar_golang_dasar/internal"
	"fmt"
)

func main(){
	fmt.Println(database.GetConnection())
}
