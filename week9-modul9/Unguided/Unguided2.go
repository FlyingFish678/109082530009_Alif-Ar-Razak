package main

import (
	"fmt"
	"math"
)

const NMAX = 100

type arrInt [NMAX]int

func inputArray(a *arrInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
}

func tampilSemua(a arrInt, n int) {
	fmt.Print("Semua elemen: ")
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func tampilGanjil(a arrInt, n int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func tampilGenap(a arrInt, n int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func tampilKelipatanX(a arrInt, n, x int) {
	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndeks(a *arrInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		a[i] = a[i+1]
	}
	*n--
}

func rataRata(a arrInt, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += a[i]
	}
	return float64(total) / float64(n)
}

func standarDeviasi(a arrInt, n int) float64 {
	rata := rataRata(a, n)
	variansi := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(a[i]) - rata
		variansi += selisih * selisih
	}
	return math.Sqrt(variansi / float64(n))
}

func frekuensi(a arrInt, n, bil int) int {
	count := 0
	for i := 0; i < n; i++ {
		if a[i] == bil {
			count++
		}
	}
	return count
}

func main() {
	var a arrInt
	var n int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan elemen array: ")
	inputArray(&a, n)

	tampilSemua(a, n)

	tampilGanjil(a, n)

	tampilGenap(a, n)

	var x int
	fmt.Print("Masukkan nilai x (kelipatan): ")
	fmt.Scan(&x)
	tampilKelipatanX(a, n, x)

	var idx int
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&idx)
	hapusIndeks(&a, &n, idx)
	fmt.Print("Setelah hapus indeks ", idx, ": ")
	tampilSemua(a, n)

	fmt.Printf("Rata-rata: %.2f\n", rataRata(a, n))

	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(a, n))

	var bil int
	fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&bil)
	fmt.Printf("Frekuensi %d: %d\n", bil, frekuensi(a, n, bil))
}
