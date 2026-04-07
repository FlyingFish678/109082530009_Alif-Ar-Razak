package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14

func volume(t, r float64) float64 {
	return pi * r * r * t
}

func massa(t, r, p float64) float64 {
	return volume(t, r) * p
}

func display(m1, m2 float64) {
	const eps = 1e-9

	if math.Abs(m1-m2) < eps {
		fmt.Println("BALANCE")
	} else {
		fmt.Printf("Selisih massa zat cair kiri dan massa zat cair kanan : %g\n", math.Abs(m1-m2))
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mJKiri, mJKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mJKiri)

	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mJKanan)

	massaKiri = massa(tKiri, r, mJKiri)
	massaKanan = massa(tKanan, r, mJKanan)

	display(massaKiri, massaKanan)
}
