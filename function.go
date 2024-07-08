package main

import "fmt"


type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func sayHello() {
	println("Hello World")
}

func sayHelloTo(name string) {
	println("Hello", name)
}

func sum(a int, b int) int {
	return a + b
}

func getSumMultiply(a, b int) (int, int) {
	return a + b, a * b
}

func getFullName(firstName, middleName, lastName string)(firstNameResult, middleNameResult, lastNameResult string) {
	firstNameResult = firstName
	middleNameResult = middleName
	lastNameResult = lastName

	return
}

func sumAll(numbers ...int) int{
	total := 0;
	for _,number := range numbers {
		total += number
	}

	return total
}

func getName(name string) string {
	return name
}

func sayHelloWithFilter(name string, filter func(string)string){
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func logging(){
	message := recover()
	if message != nil{
		fmt.Println("Error dengan message", message)
	
	}
	fmt.Println("Selesai memanggil function")
}

func runApplication(){
	defer logging()
	fmt.Println("Run Application")
}


func runApp(error bool){
	defer logging()

	if error {
		panic("APLIKASI ERROR")
	}
}



func main() {
	sayHello()

	sayHelloTo("M Bimo Bayu Bagaskara")

	fmt.Println(sum(10, 10))

	result1, result2 := getSumMultiply(10, 10)

	fmt.Println(result1)
	fmt.Println(result2)

	result3, _ := getSumMultiply(10, 10)

	fmt.Println(result3)

	firstName, middleName, lastName := getFullName("M", "Bimo", "Bayu Bagaskara")

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fmt.Println(sumAll(1, 2, 3, 4, 5))

	name:= getName

	fmt.Println(name("M Bimo Bayu Bagaskara"))

	sayHelloWithFilter("anjing", spamFilter)

	sayHelloWithFilter("bimo", spamFilter)

	blacklist := func(name string) bool {
		if name == "anjing" {
			return true
		} else {
			return false
		}
	}

	registerUser("anjing", blacklist)

	registerUser("bimo", blacklist)

	runApplication()

	runApp(true)
}