package main

import "fmt"

const NMAX = 100

type arrString [NMAX]string

func main() {
	var klubA, klubB string
	var hasil arrString
	var n int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}

		fmt.Printf("Hasil %d : %s\n", pertandingan, hasil[n])
		n++
		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
}
