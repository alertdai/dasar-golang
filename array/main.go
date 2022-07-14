package main

import "fmt"

func main()  {
	var names [3]string
	names[0] = "Not Good,"
	names[1] = "Not Bad,"
	names[2] = "Just Fine"
	fmt.Println(names[0], names[1], names[2])

	fmt.Println()
	fmt.Println("Inisialisasi Nilai Awal Array")
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Isi semua element \t:", fruits)
	fmt.Println("Jumlah element \t:", len(fruits))
	fmt.Println("ukuran array \t:", cap(fruits))

	fmt.Println()
	fmt.Println("Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen")
	var numbers = [...]int{2,3,2,4,3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))
	fmt.Println("ukuran array \t:", cap(numbers))

	fmt.Println()
	fmt.Println("Array Multidimensi")
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	fmt.Println()
	fmt.Println("Perulangan Elemen Array Menggunakan Keyword for")
	var cars = [4]string{"honda", "suzuki", "hyundai", "bmw"}
	for i := 0; i < len(cars); i++ {
		fmt.Printf("elemen %d: %s\n", i, cars[i])
	}

	fmt.Println()
	fmt.Println("Perulangan Elemen Array Menggunakan Keyword for - range")
	for i, cars := range cars {
		fmt.Printf("elemen %d: %s\n", i, cars)
	}

	fmt.Println()
	fmt.Println("Penggunaan Variabel Underscore _ Dalam for - range")
	for _, cars := range cars {
		fmt.Printf("merk mobil : %s\n", cars)
	}

	//Jika yang dibutuhkan hanya indeks elemen-nya saja, bisa gunakan 1 buah variabel setelah keyword for.
	fmt.Println()
	for i, _ := range cars { // atau bisa dengan cara = for i := range fruits { }
		fmt.Printf("indexnya : %d\n", i)
	}

	fmt.Println()
	var motor = make([]string,2)
	motor[0] = "Honda"
	motor[1] = "Yamaha"
	fmt.Print(motor)
}