package main
import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int = 0
	var suaraMasuk int

	fmt.Println("masukkan proses input suara :")
	fmt.Scan(&suaraMasuk)
	for suaraMasuk != -1 {
		idx := posisi(p, n, suaraMasuk)
		if idx != -1 {
			p[idx].suara++
		} else {
			p[n].nama = suaraMasuk
			p[n].suara = 1
			n++
		}
		fmt.Scan(&suaraMasuk)
	}

	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}

	fmt.Println()
	fmt.Println("hasil perhitungan suara: ")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
	}
	fmt.Println()
}