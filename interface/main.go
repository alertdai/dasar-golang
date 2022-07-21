package main

import (
	lib "Belajar_Golang/dasar-golang/interface/library"
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func pemisah(judul string) {
	fmt.Println()
	fmt.Println(judul)
}

func main() {
	var judul string

	judul = "Penerapan Interface"
	pemisah(judul)
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("==== Persegi ====")
	fmt.Println("Luas  \t\t:", bangunDatar.luas())
	fmt.Println("Keliling  \t:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("==== Lingkaran ====")
	fmt.Println("Luas  \t\t:", bangunDatar.luas())
	fmt.Println("Keliling  \t:", bangunDatar.keliling())
	fmt.Println("Jari - Jari \t:", bangunDatar.(lingkaran).jariJari())
	// Atau ada cara casting Method jariJari() dengan cara dibawah ini

	var bangunLingkaran lingkaran = bangunDatar.(lingkaran)
	fmt.Println("==== Akses jariJari() ====")
	fmt.Println("Jari - Jari \t:", bangunLingkaran.jariJari())

	judul = "Embedded Interface"
	pemisah(judul)
	var bangunRuang lib.NewHitung = &lib.Kubus{4}
	fmt.Println("==== Kubus ====")
	fmt.Println("Luas  \t\t:", bangunRuang.Luas())
	fmt.Println("Keliling  \t:", bangunRuang.Keliling())
	fmt.Println("Volume \t\t:", bangunRuang.Volume())
}
