package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku() {
	fmt.Scan(&nPustaka)
	var i int
	i = 0
	for i < nPustaka {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
		i = i + 1
	}
}

func CetakTerfavorit() {
	idx := 0
	var i int
	i = 1
	for i < nPustaka {
		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
		i = i + 1
	}
	fmt.Println(pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun)
}

func UrutBuku() {
	var temp Buku
	var i, j int
	i = 1
	for i <= nPustaka-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru() {
	limit := 5
	if nPustaka < 5 {
		limit = nPustaka
	}
	var i int
	i = 0
	for i < limit {
		fmt.Println(pustaka[i].judul)
		i = i + 1
	}
}

func CariBuku(r int) {
	low := 0
	high := nPustaka - 1
	ketemu := -1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			ketemu = mid
			break
		} else if pustaka[mid].rating > r {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[ketemu]
		fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	}
}

func main() {
	DaftarkanBuku()
	CetakTerfavorit()
	UrutBuku()
	Cetak5Terbaru()
	var r int
	fmt.Scan(&r)
	CariBuku(r)
}
