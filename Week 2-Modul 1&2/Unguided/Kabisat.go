package main
import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)
	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		fmt.Println(tahun, "adalah tahun kabisat")
	} else {
		fmt.Println(tahun, "bukan tahun kabisat")
	}
}