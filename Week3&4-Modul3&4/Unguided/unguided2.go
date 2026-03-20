package main

import (
	"fmt"
)

func main() {

	var x, y, z int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scanln(&y)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scanln(&z)

	f := func(val int) int { return val * val }
	g := func(val int) int { return val - 2 }
	h := func(val int) int { return val + 1 }

	fmt.Printf("F(G(H( %d ))) : %d\n", x, f(g(h(x))))
	fmt.Printf("G(H(F( %d ))) : %d\n", y, g(h(f(y))))
	fmt.Printf("H(F(G( %d ))) : %d\n", z, h(f(g(z))))
}
