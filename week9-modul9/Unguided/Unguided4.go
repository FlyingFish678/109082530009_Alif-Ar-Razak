package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	fmt.Printf("Hves\t: ")
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != ' ' && ch != '\n' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	fmt.Printf("Rvvvusv tves : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel
	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}
	balikanArray(&balik, n)
	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	cetakArray(tab, m)

	fmt.Printf("flalitauou\t: o %v\n", palindrom(tab, m))

	fmt.Println()

	balikanArray(&tab, m)

	cetakArray(tab, m)
}
