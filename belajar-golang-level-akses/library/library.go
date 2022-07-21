package library

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func Pemisah(judul string) {
	fmt.Println()
	fmt.Println(judul)
}

func SayHello(name string) {
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("My name is", name)
}

var Personal = struct {
	Name string
	Age  int
}{}

func init() {
	Personal.Name = "Rendi Santosa"
	Personal.Age = 25

	fmt.Println("--> library/library.go imported")
}
