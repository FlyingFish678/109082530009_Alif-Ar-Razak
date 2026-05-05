package main

import (
	"fmt"
	"math"
)

const nMax int = 100

type NamaProv [nMax]string
type PopProv [nMax]int
type TumbuhProv [nMax]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv, n int) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func provinsiTercepat(tumbuh TumbuhProv, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

func prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv, n int) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < n; i++ {
		if tumbuh[i] > 0.02 {
			prediksiPop := math.Round((tumbuh[i] + 1) * float64(pop[i]))
			fmt.Printf("%s %d\n", prov[i], int(prediksiPop))
		}
	}
}

func indeksProvinsi(prov NamaProv, n int, nama string) int {
	for i := 0; i < n; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var n int

	fmt.Print("Masukkan jumlah provinsi: ")
	fmt.Scan(&n)

	inputData(&prov, &pop, &tumbuh, n)

	var namaCari string
	fmt.Print("Masukkan nama provinsi yang dicari: ")
	fmt.Scan(&namaCari)

	idx := provinsiTercepat(tumbuh, n)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idx])

	idxCari := indeksProvinsi(prov, n, namaCari)
	fmt.Printf("\nData provinsi yang dicari : %s (Index: %d)\n", namaCari, idxCari)

	prediksi(prov, pop, tumbuh, n)
}
