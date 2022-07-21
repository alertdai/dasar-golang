package main

import (
	"fmt"
	"strings"
)

func pemisah(judul string) {
	fmt.Println()
	fmt.Println(judul)
}

type person struct {
	name string
	age  int
}

func main() {
	var judul string
	judul = "Penggunaan interface{}"

	pemisah(judul)

	var secret interface{}

	secret = "Ethan Hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// Cara lain mengakses interface
	fmt.Println()
	var personal map[string]interface{}

	personal = map[string]interface{}{
		"name":  "Aldo Trapani",
		"age":   24,
		"motor": []string{"Honda", "Yamaha", "Kawasaki"},
	}

	for key, p := range personal {
		fmt.Println(key, "\t:", p)
	}

	judul = "Type Alias Any"
	pemisah(judul)

	personal = map[string]any{
		"name":  "Alex Varhoud",
		"age":   27,
		"mobil": []string{"Honda", "Suzuki", "Nissan"},
	}

	for key, p := range personal {
		fmt.Println(key, "\t:", p)
	}

	judul = "Casting Variabel Interface Kosong"
	pemisah(judul)

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multipled by 10 is", number)

	secret = []string{"A", "B", "C"}
	var alphabet = strings.Join(secret.([]string), ", ")
	fmt.Println(alphabet, "is three alphabet")

	judul = "Casting Variabel Interface Kosong Ke Objek Pointer"
	pemisah(judul)

	var secret2 interface{} = &person{name: "Wick", age: 24}
	var name = secret2.(*person).name
	fmt.Println(name)

	judul = "Kombinasi Slice, map, dan interface{}"
	pemisah(judul)

	fmt.Println("New Data 1")
	var newData1 = []map[string]interface{}{
		{"name": "Ana", "age": 25},
		{"name": "Ani", "age": 26},
		{"name": "Anu", "age": 27},
	}

	for _, each := range newData1 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	fmt.Println()
	fmt.Println("New Data 2")
	var newData2 = []interface{}{
		map[string]interface{}{"name": "Nanang", "age": 30},
		[]int{1, 2, 3},
		"orange",
	}

	for _, each := range newData2 {
		fmt.Println(each)
	}
}
