package main

import (
	// "belajar-golang-level-akses/library" // Judul ke 1 & 2
	lib "belajar-golang-level-akses/library" //Import Dengan Prefix Tanda Titik
	f "fmt"
)

func main() {
	var judul string

	// judul = "Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported"
	// library.Pemisah(judul)
	// var name string = "Wiki"
	// library.SayHello(name)

	// judul = "Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya"
	// library.Pemisah(judul)
	// var s1 = library.Student{"Idoy", 21}
	// fmt.Println("Nama:", s1.Name)
	// fmt.Println("Grade:", s1.Grade)

	judul = "Import Dengan Prefix Tanda Titik dan Pemanfaatan Alias Ketika Import Package"
	lib.Pemisah(judul)
	var s1 = lib.Student{"Jeje", 22}
	f.Println("Name:", s1.Name, "; Grade:", s1.Grade)

	judul = "Mengakses Properti Dalam File Yang Package-nya Sama"
	lib.Pemisah(judul)
	HelloWorld() // Cara menjalankannya = go run main.go partial.go

	judul = "Fungsi init()"
	lib.Pemisah(judul)
	f.Printf("Name : %s\n", lib.Personal.Name)
	f.Printf("Age : %d\n", lib.Personal.Age)
}
