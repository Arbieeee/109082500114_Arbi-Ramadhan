package main

import "fmt"

const NMAX = 1001 

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func SelectionSort(T *arrPemain, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if T[j].gol > T[idxMax].gol {
				idxMax = j
			} else if T[j].gol == T[idxMax].gol {
				if T[j].assist > T[idxMax].assist {
					idxMax = j
				}
			}
		}
		T[i], T[idxMax] = T[idxMax], T[i]
	}
}

func main() {
	var data arrPemain
	var n int

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var firstName, lastName string
		fmt.Scan(&firstName, &lastName, &data[i].gol, &data[i].assist)
		data[i].nama = firstName + " " + lastName
	}

	SelectionSort(&data, n)

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", data[i].nama, data[i].gol, data[i].assist)
	}
}