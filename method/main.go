package main

import "fmt"
import "strings"

func pemisah(judul string)  {
	fmt.Println()
	fmt.Println(judul)
}

type person struct {
	name string
	age int
}

func (p person) sayHello()  {
	fmt.Println("Halo", p.name)
}

func (p person) getNameIt(i int) string {
	return strings.Split(p.name, " ")[i-1]
}

//Method Pointer
func (p person) changeName1(name string)  {
	fmt.Println("---> on changeName1, name changed to:", name)
	p.name = name
}

func (p *person) changeName2(name string, age int)  {
	fmt.Println("---> on changeName2, name changed to:", name)
	fmt.Println("---> on changeName2, age changed to:", age)
	p.name = name
	p.age = age
}

func main()  {
	var judul string
	judul = "Penerapan Method"
	pemisah(judul)
	var p1 = person{"Alva Latif Ramdana", 24}
	p1.sayHello()
	var call = p1.getNameIt(1)
	fmt.Println("Nama panggilan:", call)

	judul = "Method Pointer"
	pemisah(judul)
	p2 := person{"Aldo Trapani", 27}
	fmt.Println("p2 name before:", p2.name)
	fmt.Println("p2 age before:", p2.age)

	p2.changeName1("Alex Varhoud")
	fmt.Println("p2 after changeName1:", p2.name)
	fmt.Println("p2 after changeName1:", p2.age)

	p2.changeName2("Sajin Komamura", 40)
	fmt.Println("p2 after changeName2:", p2.name)
	fmt.Println("p2 after changeName2:", p2.age)
	
	// Untuk codingan DI ATAS ini hanya menggati nama saja, tidak dengan umur
}