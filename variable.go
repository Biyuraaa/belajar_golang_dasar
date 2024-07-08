package main

import "fmt"

func main() {
	var name string;

	name = "M Bimo Bayu Bagaskara";

	fmt.Println(name);

	var nim = "187221049";

	fmt.Println(nim);

	jurusan := "Sistem Informasi";

	fmt.Println(jurusan);

	var (
		alamat = "Jl. Demak Jaya 2 No. 82"
		kota = "Surabaya"
		provinsi = "Jawa Timur"
	)

	fmt.Println(alamat);
	fmt.Println(kota);
	fmt.Println(provinsi);

	const gender = "Laki-laki";

	fmt.Println(gender);

	fmt.Println(string(name[0]))
}