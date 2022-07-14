package main

import "fmt"
import "strings"

func pemisah(judul string) {
	fmt.Println()
	fmt.Println(judul)
}

func firstCalculate(numbers ...int) float64 {
	const judul = "Penerapan Fungsi Variadic"
	pemisah(judul)

	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string)  {
	const judul = "Fungsi Dengan Parameter Biasa & Variadic"
	pemisah(judul)
	var hobbiesAsString = strings.Join(hobbies, ",")

	fmt.Printf("Hello my name is : %s\n", name)
	fmt.Printf("My hobbies are : %s\n", hobbiesAsString)
}

func main()  {
	var avg = firstCalculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    fmt.Println(msg)

	var judul string
	judul = "Pengisian Parameter Fungsi Variadic Menggunakan Data Slice"
	pemisah(judul)
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg2 = firstCalculate(numbers...)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)

	// Atau bisa dengan cara
	// var avg3 = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)

	fmt.Println(msg2)

	// Untuk penggunaan fungsi dengan param biasa & variadic bisa menggunakan
	// yourHobbies("wick", "coding", "football")
	// Atau
	var hobbies = []string{"coding", "football", "sleep"}
	yourHobbies("Wick", hobbies...)
}