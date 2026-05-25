package main
import "fmt"

type Pemain struct {
	namaDepan, namaBelakang string
	gol, assist             int
}

func main() {
	var n int

	fmt.Println("masukkan data input: ")
	fmt.Scan(&n)

	var arrPemain [1005]Pemain

	for i := 0; i < n; i++ {
		fmt.Scan(&arrPemain[i].namaDepan, &arrPemain[i].namaBelakang, &arrPemain[i].gol, &arrPemain[i].assist)
	}

	for i := 1; i < n; i++ {
		temp := arrPemain[i]
		j := i - 1
		
		for j >= 0 && (arrPemain[j].gol < temp.gol || (arrPemain[j].gol == temp.gol && arrPemain[j].assist < temp.assist)) {
			arrPemain[j+1] = arrPemain[j]
			j--
		}
		arrPemain[j+1] = temp
	}

	fmt.Println()

	fmt.Println("hasil sorting: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", arrPemain[i].namaDepan, arrPemain[i].namaBelakang, arrPemain[i].gol, arrPemain[i].assist)
	}
}