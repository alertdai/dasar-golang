package main

import "fmt"
import "math"

func pemisah(judul string) {
	fmt.Println()
	fmt.Println(judul)
}

func calculate(d float64) (float64, float64) {
	var judul = "Penerapan Fungsi Multiple Return"
	pemisah(judul)
	// Hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// Hitung keliling
	var circumference = math.Pi * d

	// Kembalikan 2 return
	return area, circumference
}

func simpleCalculate(e int) (nonComma int, comma float64) {
	var judul = "Fungsi Dengan Predefined Return Value"
	pemisah(judul)
	// Hitung nonComma
	nonComma = e/2
	comma = float64(e)/2

	return
}

func main()  {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling lingkaran\t: %.2f \n", circumference)

	var numbers = 3
	nonComma, comma := simpleCalculate(numbers)

	fmt.Printf("Non Comma\t: %d \n", nonComma)
	fmt.Printf("Comma\t\t: %.2f \n", comma)
}