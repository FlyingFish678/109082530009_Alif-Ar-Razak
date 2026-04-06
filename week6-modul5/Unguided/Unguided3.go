package main

import "fmt"

func ganjil(i, n int, pertama bool) {
	if i > n {
		return
	}

	if i%2 != 0 {
		if !pertama {
			fmt.Print(" ")
		}
		fmt.Print(i)
		pertama = false
	}

	ganjil(i+1, n, pertama)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	ganjil(1, n, true)
	fmt.Println()
}
