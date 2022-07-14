package main

// import "fmt"
// import "strings"

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func pemisah(judul string) {
	fmt.Println()
	fmt.Println(judul)
}

func printMessage(message string, arr []string)  {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var judul = "Fungsi Dengan Return Value / Nilai Balik"
	pemisah(judul)
	var value = rand.Int() % (max - min + 1) + min
	return value
}

func divideNumber(m, n int)  {
	var judul = "Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi"
	pemisah(judul)

	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divide by %d\n", m, n)
		return
	}

	var res = m/n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

//Catatan
/* Deklarasi Parameter Bertipe Data Sama
func nameOfFunc(paramA type, paramB type, paramC type) returnType
func nameOfFunc(paramA, paramB, paramC type) returnType

func randomWithRange(min int, max int) int
func randomWithRange(min, max int) int
*/

func main()  {
	fmt.Println("Penerapan Fungsi")
	var names = []string{"Aldo", "Trapani"}
	printMessage("Halo", names)

	rand.Seed(time.Now().Unix())
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(10,2)
	divideNumber(4,0)
}