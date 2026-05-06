package main

import "fmt"

func main() {
	const kapasitas = 1000
	var N int
	var berat [kapasitas]float64

	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&N)

	if N > kapasitas {
		N = kapasitas
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	minBerat := berat[0]
	maxBerat := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}

	fmt.Println("Data Berat Anak Kelinci")
	fmt.Printf("Berat terbesar : %.2f\n", maxBerat)
	fmt.Printf("Berat terkecil : %.2f\n", minBerat)
}