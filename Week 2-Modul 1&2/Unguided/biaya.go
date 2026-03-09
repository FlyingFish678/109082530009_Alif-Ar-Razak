package main

import "fmt"

func main() {
	var totalBerat int
	var kg, sisa int
	var biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalBerat)

	kg = totalBerat / 1000
	sisa = totalBerat % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}
	totalBiaya = biayaKg + biayaSisa
	fmt.Printf("Detail berat : %d kg + %d gram\n", kg, sisa)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}