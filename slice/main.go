package main

import "fmt"

func main()  {
	fmt.Println("Inisialisasi Slice")
	fmt.Println("Tidak perlu mengisi jumlah elemen, tidak seperti array")
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits) //atau bisa diprint dengan fruits[0] atau fruits[:]

	fmt.Println()
	fmt.Println("Operasi Slice")
	fmt.Println(fruits[1:3])
	fmt.Println(fruits[:2])
	fmt.Println(fruits[2:])

	fmt.Println()
	fmt.Println("Slice Merupakan Tipe Data Reference")
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	
	var aaFruits = aFruits[1:2]
	var bbFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	fmt.Println("Buah \"grape\" diubah menjadi \"pinnaple\"")
	bbFruits[0] = "pineapple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	fmt.Println()
	fmt.Println("Fungsi len() dan cap()")
	fmt.Println("Cek variabel \"fruits\"")
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println("Cek variabel \"cFruits\"")
	var cFruits = fruits[0:3]
	fmt.Println(len(cFruits))
	fmt.Println(cap(cFruits))

	fmt.Println("Cek variabel \"dFruits\"")
	var dFruits = fruits[1:4]
	fmt.Println(len(dFruits))
	fmt.Println(cap(dFruits))

	fmt.Println()
	fmt.Println("Fungsi append()")
	fmt.Println("Slice yang sudah dipakai fungsi append(), maka akan menjadi data referensi baru")
	var cars = []string{"Honda", "Suzuki", "Hyundai"}
	var aCars = append(cars, "Nissan")
	fmt.Println(cars)
	fmt.Println(aCars)

	fmt.Println()
	fmt.Println("Fungsi copy()")
	dst := make([]string,3)
	src := []string{"watermelon", "pineapple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
	
	fmt.Println("dst sudah diisi data")
	a := []string{"1", "2", "3"}
	b := []string{"4","5"}
	fmt.Printf("Data \"a\" belum diubah : %v\n", a)
	c := copy(a,b)
	fmt.Println("Data \"a\" sudah diubah : ",a)
	fmt.Println(b)
	fmt.Println("Data \"c\" merupakan jumlah data yang dicopy: ", c)

	fmt.Println()
	fmt.Println("Pengaksesan Elemen Slice Dengan 3 Indeks")
	var motor = []string{"honda", "yamaha", "suzuki"}
	var aMotor = motor[0:2]
	var bMotor = motor[0:2:2]

	fmt.Println("Cek variabel \"motor\"")
	fmt.Println(motor)
	fmt.Println(len(motor))
	fmt.Println(cap(motor))

	fmt.Println("Cek variabel \"aMotor\"")
	fmt.Println(aMotor)
	fmt.Println(len(aMotor))
	fmt.Println(cap(aMotor))

	fmt.Println("Cek variabel \"bMotor\"")
	fmt.Println(bMotor)
	fmt.Println(len(bMotor))
	fmt.Println(cap(bMotor))
	fmt.Println("Saat menggunakan slice dengan 3 indeks, size pada \"bMotor\" akan berubah menjadi 2")
	fmt.Println("Karena saat menggunakan slice 3 indeksnya menggunakan [0:2:2] (yang paling kanan)")
}