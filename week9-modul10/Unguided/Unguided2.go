package main

import (
	"fmt"
)

type arrIkan [1000]float64

func main() {
	var x, y int
	var ikan arrIkan

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalSemuaBerat float64 = 0
	var beratPerWadah float64 = 0
	var hitungIkan int = 0
	var jumlahWadah int = 0

	for i := 0; i < x; i++ {
		beratPerWadah += ikan[i]
		totalSemuaBerat += ikan[i]
		hitungIkan++

		if hitungIkan == y || i == x-1 {
			fmt.Printf("%.2f ", beratPerWadah)
			beratPerWadah = 0
			hitungIkan = 0
			jumlahWadah++
		}
	}
	fmt.Println()

	if jumlahWadah > 0 {
		rataRata := totalSemuaBerat / float64(jumlahWadah)
		fmt.Printf("%.2f\n", rataRata)
	}
}