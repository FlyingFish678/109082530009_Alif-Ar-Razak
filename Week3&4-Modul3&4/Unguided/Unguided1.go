package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int

	fmt.Print("masukkan nilai a : ")
	fmt.Scanln(&a)
	fmt.Print("masukkan nilai b : ")
	fmt.Scanln(&b)
	fmt.Print("masukkan nilai c : ")
	fmt.Scanln(&c)
	fmt.Print("masukkan nilai d : ")
	fmt.Scanln(&d)

	var fak func(n int) int
	fak = func(n int) int {
		if n <= 1 {
			return 1
		}
		return n * fak(n-1)
	}

	perAC := fak(a) / fak(a-c)
	komAC := perAC / fak(c)

	perBD := fak(b) / fak(b-d)
	komBD := perBD / fak(d)

	fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", a, c, perAC)
	fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", a, c, komAC)
	fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", b, d, perBD)
	fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", b, d, komBD)
}
