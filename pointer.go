package main

import "fmt"

type MahasiswaSI struct {
	name, nim, jurusan string
	ipk                int
	beasiswa bool
}


func changeJurusanToTeknikInformatika(mahasiswa *MahasiswaSI){
	mahasiswa.jurusan = "Teknik Informatika"
}

func (mahasiswa *MahasiswaSI) gotBeasiswa (){
	if mahasiswa.ipk > 3 {
		mahasiswa.beasiswa = true	
	}
}


func main() {
	//Pass by value
	bimo := MahasiswaSI{"M Bimo Bayu Bagaskara", "187221049", "Sistem Informasi", 3,false}
	bimo2 := bimo

	bimo.name = "Biyuraaaa"

	fmt.Println(bimo2.name)

	//Gunakan operator & untuk melakukan pass by reference
	bimo3 := &bimo

	bimo3.name = "Biyuraaa"

	fmt.Println(bimo.name)

	//Cara mengubah dari pass by reference  ke address baru menggunakan asterisk operator (*)
	
	bimo3 = &MahasiswaSI{"M Bimo Bayu Bagaskara", "187221049", "Sistem Informasi", 3,false}

	bimo3.ipk = 4

	fmt.Println(bimo.ipk)
	fmt.Println(bimo3.ipk)

	//Menggunakan asterisk (*) sebelum variable untuk mengubah seluruh variabel yang referensi ke address tersebut ke address baru
	*bimo3 = MahasiswaSI{"M Bimo Bayu Bagaskara", "187221049", "Sistem Informasi", 3,false}

	changeJurusanToTeknikInformatika(bimo3)

	fmt.Println(bimo3.jurusan)
	fmt.Println(bimo.jurusan)

	bimo3.ipk = 4
	fmt.Println(bimo.beasiswa)
	bimo3.gotBeasiswa()

	fmt.Println(bimo.beasiswa)

}


