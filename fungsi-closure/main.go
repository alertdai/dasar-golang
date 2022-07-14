package main

import "fmt"

func pemisah(judul string)  {
	fmt.Println()
	fmt.Println(judul)
}

func main()  {
	var judul string
	judul = "Closure Disimpan Sebagai Variabel"
	pemisah(judul)

	var getMinMax = func (n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max :
				max = e
			case e < min :
				min = e
			}
		}
		return min, max
	}

	var numbers []int
	numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf(" Data: %v\n Min: %v\n Max: %v\n", numbers, min, max)

	judul = "Immediately-Invoked Function Expression (IIFE)"
	pemisah(judul)
	
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	
	var newNumbers = func(min int) []int{
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3) // (3) adalah ciri khas dari IIFE

	fmt.Println("Original number: ", numbers)
	fmt.Println("Filtered number:", newNumbers)

	judul = "Closure Sebagai Nilai Kembalian"
	pemisah(judul)

	var numMax = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, numMax)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int)  {
	var res []int
	for _, e := range numbers {
		if e < max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
	// Fungsi tanpa nama pun bisa disimpan ke dalam variabel dulu dengen contoh :
	/*
	var getNumbers = func() []int {
		return res
	}
	return len(res), getNumbers
	*/
}