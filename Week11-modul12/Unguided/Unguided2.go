package main

import "fmt"

func main() {
	var suara [21]int
	total := 0
	sah := 0

	var x int
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		total++
		if x >= 1 && x <= 20 {
			sah++
			suara[x]++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)

	max1, max2 := 0, 0
	ketua, wakil := -1, -1

	for i := 1; i <= 20; i++ {
		if suara[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = suara[i]
			ketua = i
		} else if suara[i] > max2 {
			max2 = suara[i]
			wakil = i
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
