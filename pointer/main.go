package main

import "fmt"

func pemisah(judul string)  {
	fmt.Println()
	fmt.Println(judul)
}

func main()  {
	var judul string
	var valNul *int
	var valNotNull int

	fmt.Println(valNul)
	fmt.Println(valNotNull)
	fmt.Println(judul)

	judul = "Penerapan Pointer"
	pemisah(judul)
	
	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	judul = "Efek Perubahan Nilai Pointer"
	pemisah(judul)

	numberA = 4

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	judul = "Parameter Pointer"
	pemisah(judul)
	var number = 10
	fmt.Println("Before change :", number)

	change(&number, 15)
	fmt.Println("After change :", number)
}

func change(original *int, value int)  {
	*original = value
}