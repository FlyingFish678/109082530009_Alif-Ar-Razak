package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	p := posisi(n, k)
	if p == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(p)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1
	found := -1
	for kr <= kn && found == -1 {
		med := (kr + kn) / 2
		if k < data[med] {
			kn = med - 1
		} else if k > data[med] {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}
