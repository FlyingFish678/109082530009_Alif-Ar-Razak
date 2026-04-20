package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var dataBuku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	fmt.Scan(&dataBuku.judul)

	fmt.Print("Masukkan nama penulis : ")
	fmt.Scan(&dataBuku.penulis)

	fmt.Print("Masukkan penerbit : ")
	fmt.Scan(&dataBuku.penerbit)

	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scan(&dataBuku.tahunTerbit)

	fmt.Print("Masukkan jumlah halaman : ")
	fmt.Scan(&dataBuku.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", dataBuku.judul)
	fmt.Println("Penulis :", dataBuku.penulis)
	fmt.Println("Penerbit :", dataBuku.penerbit)
	fmt.Println("Tahun Terbit :", dataBuku.tahunTerbit)
	fmt.Println("Jumlah Halaman :", dataBuku.jumlahHalaman)
}
