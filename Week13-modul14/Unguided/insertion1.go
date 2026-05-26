package main

import "fmt"

type arrInt [1000000]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var arr arrInt
	var n int
	n = 0

	var x int
	fmt.Scan(&x)
	for x >= 0 {
		arr[n] = x
		n = n + 1
		fmt.Scan(&x)
	}

	insertionSort(&arr, n)

	var i int
	i = 0
	for i < n {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
		i = i + 1
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := arr[1] - arr[0]
	sama := true
	i = 2
	for i < n {
		if arr[i]-arr[i-1] != jarak {
			sama = false
		}
		i = i + 1
	}

	if sama {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
