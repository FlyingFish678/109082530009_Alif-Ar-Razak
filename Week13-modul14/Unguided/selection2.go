package main

import "fmt"

type arrInt [1000000]int

func selectionSortAsc(T *arrInt, n int) {
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

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
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
		var ganjil, genap arrInt
		var ng, nge int
		ng = 0
		nge = 0
		var j int
		j = 0
		for j < m {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil[ng] = x
				ng = ng + 1
			} else {
				genap[nge] = x
				nge = nge + 1
			}
			j = j + 1
		}

		selectionSortAsc(&ganjil, ng)
		selectionSortDesc(&genap, nge)

		first := true
		var k int
		k = 0
		for k < ng {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[k])
			first = false
			k = k + 1
		}
		k = 0
		for k < nge {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(genap[k])
			first = false
			k = k + 1
		}
		fmt.Println()
		i = i + 1
	}
}
