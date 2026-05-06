package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var arrBerat arrBalita
	var n int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	if n > 100 {
		n = 100
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	var minBerat, maxBerat float64
	hitungMinMax(arrBerat, n, &minBerat, &maxBerat)

	rataRata := rerata(arrBerat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", minBerat)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxBerat)
	fmt.Printf("Berat balita rata-rata: %.2f kg\n", rataRata)
}