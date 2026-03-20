package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scanln(&n)

	cetakDeret := func(n int) {
		for {
			fmt.Print(n)

			if n == 1 {
				fmt.Println()
				break
			}

			fmt.Print(" ")

			if n%2 == 0 {
				n = n / 2
			} else {
				n = (3 * n) + 1
			}
		}
	}

	cetakDeret(n)
}
