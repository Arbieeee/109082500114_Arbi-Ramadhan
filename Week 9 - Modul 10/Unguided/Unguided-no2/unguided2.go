package main

import "fmt"

func main() {
	const kapasitas = 1000
	var x, y int
	var berat [kapasitas]float64

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x > kapasitas {
		x = kapasitas
	}

	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y

	totalPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalPerWadah[indeksWadah] += berat[i]
	}

	fmt.Println("\nTotal berat ikan di setiap wadah")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah ke-%d = %.2f\n", i+1, totalPerWadah[i])
	}

	var totalSemua float64
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalPerWadah[i]
	}
	rataRata := totalSemua / float64(jumlahWadah)

	fmt.Println("\nRata-rata berat ikan per wadah:")
	fmt.Printf("%.2f\n", rataRata)
}