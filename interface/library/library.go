package library

import "math"

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type NewHitung interface {
	Hitung2d
	Hitung3d
}

type Kubus struct {
	Sisi float64
}

func (k *Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

func (k *Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

func (k *Kubus) Keliling() float64 {
	return k.Sisi * 12
}
