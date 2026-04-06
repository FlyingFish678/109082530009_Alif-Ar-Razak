package main

import "fmt"

func faktor(n, i int, pertama bool) {
	if i > n {
		return
	}

	if n%i == 0 {
		if !pertama {
			fmt.Print(" ")
		}
		fmt.Print(i)
		pertama = false
	}

	faktor(n, i+1, pertama)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	faktor(n, 1, true)
	fmt.Println()
}
