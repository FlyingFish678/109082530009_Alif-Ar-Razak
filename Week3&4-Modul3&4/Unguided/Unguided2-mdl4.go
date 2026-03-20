package main

import (
	"fmt"
	"math"
)

func main() {
	hitungPersegi := func(sisi int) {
		fmt.Printf("Luas persegi : %d\n", sisi*sisi)
		fmt.Printf("Keliling persegi : %d\n", 4*sisi)
	}

	hitungPersegiPanjang := func(panjang, lebar int) {
		fmt.Printf("Luas persegi panjang : %d\n", panjang*lebar)
		fmt.Printf("Keliling persegi panjang : %d\n", 2*(panjang+lebar))
	}

	hitungLingkaran := func(r float64) {
		fmt.Printf("Luas lingkaran : %v\n", math.Pi*r*r)
		fmt.Printf("Keliling lingkaran : %v\n", 2*math.Pi*r)
	}

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")

	var pilihan int
	fmt.Scanln(&pilihan)
	fmt.Println()

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukkan sisi : ")
		fmt.Scanln(&sisi)
		hitungPersegi(sisi)
	case 2:
		var p, l int
		fmt.Print("Masukkan panjang : ")
		fmt.Scanln(&p)
		fmt.Print("Masukkan lebar : ")
		fmt.Scanln(&l)
		hitungPersegiPanjang(p, l)
	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scanln(&r)
		hitungLingkaran(r)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
