package main

import "fmt"

type arrInt [1000000]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	var i int
	i = 0
	for i < n {
		fmt.Scan(&m)
		var arr arrInt
		var j int
		j = 0
		for j < m {
			fmt.Scan(&arr[j])
			j = j + 1
		}
		selectionSort(&arr, m)
		var k int
		k = 0
		for k < m {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(arr[k])
			k = k + 1
		}
		fmt.Println()
		i = i + 1
	}
}
