package main

import "fmt"

func main() {
	// Perintah pertama
	fmt.Println("Hello World")

	// Go Command
	// Command 1 baris
	/*
	Command untuk banyak baris
	*/

	//Variabel
	var firstname string = "John"
	var lastname string
	lastname = "Wick"

	fmt.Printf("\nHalo %s %s!\n", firstname, lastname)

	nomor_1, nomor_2, nomor_3 := "satu", 2, 3.00

	fmt. Printf("Perbandingan angka : %s, %v, %v\n", nomor_1, nomor_2, nomor_3)
	//nama variabel yang tidak harus dideklarasikan
	_ = "Test"

	//Seleksi Kondisi
	var point int
	point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	point = 2
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	point = 6
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("just fine")
	default :
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//Perulangan
	fmt.Println("")

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("Penggunaan Keyword for Dengan Argumen Hanya Kondisi")

	var i int
	i = 0
	for i <= 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("Penggunaan Keyword for Tanpa Argumen")

	i = 0
	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {break}
	}

	fmt.Println("Pemanfaatan Label Dalam Perulangan")
	outerloop:
	for a := 0; a < 5; a++{
		for b := 0; b < 5; b++{
			if a == 3{
				break outerloop
			}
			fmt.Print("matriks [",a,"][",b,"]", "\n")
		}
	}
}
