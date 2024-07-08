package main

func main() {
	bimo := map[string]string{
		"name":    "M Bimo Bayu Bagaskara",
		"nim":     "187221049",
		"jurusan": "Sistem Informasi",
	}

	if bimo["name"] == "M Bimo Bayu Bagaskara" {
		println("Nama saya adalah M Bimo Bayu Bagaskara")
	} else {
		println("Nama saya bukan M Bimo Bayu Bagaskara")
	}

	if bimo["jurusan"] == "Sistem Informasi" {
		println("Saya Jago Coding")
	} else if bimo["jurusan"] == "Teknik Informatika" {
		println("Saya Jago Coding")
	} else {
		println("Saya tidak jago coding")
	}

	if length := len(bimo["name"]); length > 10 {
		println("Nama saya panjang")
	} else {
		println("Nama saya pendek")
	}
}