package main

import "fmt"

func pemisah(judul string)  {
	fmt.Println()
	fmt.Println(judul)
}

type student struct {
	name string
	grade int
}

type person struct {
	name string
	age int
}

type motor struct {
	motor string
	person
	age int // "Embedded Struct Dengan Nama Property Yang Sama"
}

// Deklarasi Dan Inisialisasi Struct Secara Horizontal, seperti cara dibawah
//type person struct { name string; age int; hobbies []string }

//Tag property dalam struct
	/*
	type person struct {
		name string `tag1`
		age  int    `tag2`
	}
	*/

type newPerson struct {
	name string
	age int
}
type newPeople = newPerson

func main()  {
	var judul string
	judul = "Deklarasi dan Penerapan Struct"

	pemisah(judul)
	var s1 student
	s1.name = "John"
	s1.grade = 2

	fmt.Println("name \t:", s1.name)
	fmt.Println("grade \t:", s1.grade)

	judul = "Inisialisasi Struct"
	pemisah(judul)
	var s2 = student{}
	s2.name = "Aldo"
	s2.grade = 3

	var s3 = student{"Alex", 4}

	var s4 = student{name: "Jason"}

	fmt.Println("Nama setiap student")
	fmt.Println("S2 name \t:", s2.name)
	fmt.Println("S3 name \t:", s3.name)
	fmt.Println("S4 name \t:", s4.name)

	fmt.Println()
	fmt.Println("Nilai setiap student")
	fmt.Println("S2 grade \t:", s2.grade)
	fmt.Println("S3 grade \t:", s3.grade)
	fmt.Println("S4 grade \t:", s4.grade)

	judul = "Variabel Objek Pointer"
	pemisah(judul)

	var s5 *student = &s4
	fmt.Println("student 4 name \t:", s4.name)
	fmt.Println("student 5 name \t:", s5.name)

	s5.name = "Ethan"
	fmt.Println("student 4 name \t:", s4.name)
	fmt.Println("student 5 name \t:", s5.name)

	judul = "Embedded Struct"
	pemisah(judul)
	var m1 = motor{}
	m1.name = "Nathan"
	m1.age = 24
	m1.motor = "Honda"
	fmt.Println("name \t:", m1.name)
	fmt.Println("age \t:", m1.age)
	fmt.Println("age \t:", m1.person.age)
	fmt.Println("motor \t:", m1.motor)

	judul = "Embedded Struct Dengan Nama Property Yang Sama"
	pemisah(judul)
	m1.person.age = 25
	fmt.Println("name \t:", m1.name)
	fmt.Println("age \t:", m1.person.age) // Age from struct "person"
	fmt.Println("age \t:", m1.age) // Age from struct "motor"
	fmt.Println("motor \t:", m1.motor)

	judul = "Pengisian Nilai Sub-Struct"
	pemisah(judul)
	var p2 = person{name: "Sajin", age:21}
	var m2 = motor{motor: "Yamaha", person: p2, age:22}
	fmt.Println("name \t:", m2.name)
	fmt.Println("age \t:", m2.person.age) // Age from struct "person"
	fmt.Println("age \t:", m2.age) // Age from struct "motor"
	fmt.Println("motor \t:", m2.motor)

	judul = "Anonymous Struct"
	pemisah(judul)
	var mo1 = struct {
		person
		mobil string
	}{ 
	// Bisa juga untuk mengisi data disetiap variabel dengan menggunakan cara dibawah 
	// (Didalam kurung setelah deklarasi struct anonymous)
		person: person{name: "Oga", age:27},
		mobil: "Daihatsu",
	}
	// Atau bisa juga mengisi data disetiap variabel dengan cara dibawah
	// mo1.person = person{name:"Ryu", age:25}
	// mo1.mobil = "Nissan"
	fmt.Println("name \t:", mo1.name) // Bisa juga mo1.person.name
	fmt.Println("age \t:", mo1.age) // Bisa juga mo1.person.age
	fmt.Println("mobil \t:", mo1.mobil)

	judul = "Kombinasi Slice & Struct"
	pemisah(judul)
	var allNewPersons = []person{
		{name: "Warno", age: 24},
		{name: "Edo", age: 26},
		{name: "Idoy", age: 21},
	}

	for _, newPerson := range allNewPersons {
		fmt.Println("Name:", newPerson.name, "\nAge:", newPerson.age)
	}

	judul = "Inisialisasi Slice Anonymous Struct"
	pemisah(judul)
	var allStudents = []struct { //[]struct sudah termasuk deklarasi dan inisialisasi
		person
		grade int
	}{
		{person: person{name: "Aldo", age: 21}, grade: 7},
		{person: person{name: "Alex", age: 22}, grade: 8},
		{person: person{name: "Alva", age: 23}, grade: 9},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}

	judul = "Deklarasi Anonymous Struct Menggunakan Keyword var"
	pemisah(judul)
	var otherPerson struct { // ini hanya deklarasi
		person
		city string
	}

	otherPerson.person = person{name: "Maman", age:20}
	otherPerson.city = "Bandung"

	fmt.Println(otherPerson)

	// Nested Struct
	/*
	type student struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}
	*/

	//Deklarasi Dan Inisialisasi Struct Secara Horizontal
	/*
	var p1 = struct { name string; age int } { age: 22, name: "wick" }
	var p2 = struct { name string; age int } { "ethan", 23 }
	*/

	judul = "Type Alias"
	pemisah(judul)
	var pe1 = newPeople{"Bagas", 19}
	fmt.Println(pe1)
	pe2 := newPeople{"Raka", 18}
	fmt.Println(pe2)
}