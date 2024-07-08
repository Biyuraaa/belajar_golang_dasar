package main

func main() {
	names := [...]string{"M", "Bimo", "Bayu", "Bagaskara"}

	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	for i, name := range names {
		println(i, name)
	}

	for _, name := range names {
		println(name)
	}
}