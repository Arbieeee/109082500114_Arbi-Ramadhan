package main

import "fmt"

const NMAX = 1000000 // jumlah maks data masukan

type arrInt [NMAX]int // tipe data alias array integer

func SelectionSort(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}
		T[i], T[idxMin] = T[idxMin], T[i]
	}
}

func median(T arrInt, n int) float64 {
	if n%2 == 1 {
		return float64(T[n/2])
	}
	return float64(T[n/2-1]+T[n/2]) / 2.0
}

func main() {
	var A arrInt // array integer
	var x int // variabel massukan
	var n int = 0

	fmt.Println("Input data masukan :")

	fmt.Scan(&x)

	for x != -5313541 && n < NMAX {
		if x == 0 {
			SelectionSort(&A, n)
			fmt.Printf("Median :\n%.0f\n", median(A, n))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}