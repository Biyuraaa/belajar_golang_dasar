package main

type Mahasiswa interface {
	getName() string
	getNIM() string
	getJurusan() string
}

type SistemInformasi struct {
	Name, NIM, Jurusan string
}

func (si SistemInformasi) getName() string {
	return si.Name
}

func (si SistemInformasi) getNIM() string {
	return si.NIM
}

func (si SistemInformasi) getJurusan() string {
	return si.Jurusan
}

func main() {
	bimo := SistemInformasi{
		Name:    "M Bimo Bayu Bagaskara",
		NIM:     "187221049",
		Jurusan: "Sistem Informasi",
	}

	println(bimo.getName())
	println(bimo.getNIM())
	println(bimo.getJurusan())
}