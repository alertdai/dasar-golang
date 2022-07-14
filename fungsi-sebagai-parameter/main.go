package main

import "fmt"
import "strings"

func pemisah(judul string)  {
	fmt.Println()
	fmt.Println(judul)
}

func filter(data []string, callback func(string) bool) []string {
	var judul = "Penerapan Fungsi Sebagai Parameter"
	pemisah(judul)
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered { // pada tulisan filtered terakhir bisa kita ganti kondisinya dengan filtered == false
			fmt.Println(filtered, "\t:", each)
			result = append(result, each)
		} else {
			fmt.Println(filtered, "\t:", each)
		}
	}
	return result
}

// Alias Skema Closure
// Cara membuat alias nama function yang akan dijadikan parameter di sebuah function
// type FilterCallback func(string) bool
// func filter(data []string, callback FilterCallback) []string{
// 	// Disini perintah	
// }

func main()  {
	var data = []string{"wick", "jason", "ethan", "aldo"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each)==5
	})

	fmt.Println()
	fmt.Println("data asli \t\t:", data)
	// data asli : ["wick", "jason", "ethan", "aldo"]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason, aldo]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
	// filter jumlah huruf 5 : [jason, ethan]
}