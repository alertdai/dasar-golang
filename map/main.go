package main

import "fmt"

func main()  {
	fmt.Println("Penggunaan Map")
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("februari", chicken["februari"])

	fmt.Println()
	fmt.Println("Inisialisasi Map")
	data := map[string]int{}
	data["one"] = 1
	fmt.Println(data["one"])
	fmt.Println("Inisialisasi Map cara lain")
	var cow = map[string]int{
		"januari" : 50,
		"februari" : 40,
	}
	fmt.Println(cow)
	//Cara inisialisasi map tanpa nilai awal
	/*
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int
	*/

	fmt.Println()
	fmt.Println("Iterasi Item Map Menggunakan for - range")
	var sheep = map[string]int {
		"januari":30,
		"februari":25,
		"maret":20,
		"april":15,
	}
	for key, val := range sheep {
		fmt.Println(key, "\t:", val)
	}
	fmt.Println(len(sheep))

	fmt.Println()
	fmt.Println("Menghapus Item Map")
	delete(sheep, "maret")
	for key, val := range sheep {
		fmt.Println(key, "\t:", val)
	}
	fmt.Println(len(sheep))

	fmt.Println()
	fmt.Println("Deteksi Keberadaan Item Dengan Key Tertentu")
	var pig = map[string]int{"januari":50, "februari":40}
	fmt.Println(pig)
	var value, isExist = pig["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	fmt.Println()
	fmt.Println("Kombinasi Slice & Map")
	var horses = []map[string]string{
		map[string]string{"name":"red", "gender":"male"},
		map[string]string{"name":"blue", "gender":"female"},
		map[string]string{"name":"green", "gender":"male"},
	}

	for _, horse := range horses {
		fmt.Println(horse["name"], "\t:", horse["gender"])
	}

	fmt.Println("Contoh Kombinasi yg lain dan key yg berbeda beda")
	var peoples = []map[string]string {
		{"name": "Aldo Trapani", "age": "35", "gender":"Male"},
		{"address": "Litle Italy", "id": "73219"},
		{"community": "Family Corleone"},
	}
	for i := 0; i < len(peoples); i++ {
		for key, people := range peoples[i] {
			fmt.Println(key, "    ","\t:",people)
		}
	}
}