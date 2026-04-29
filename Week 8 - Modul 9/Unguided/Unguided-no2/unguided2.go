package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	fmt.Print("Masukkan jumlah elemen N: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Printf("Masukkan %d buah bilangan bulat: ", N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println()

	fmt.Print("a. Isi seluruh array: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai X untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	var hapusIdx int
	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&hapusIdx)
	arr = append(arr[:hapusIdx], arr[hapusIdx+1:]...)
	fmt.Print("e. Isi array setelah dihapus: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	rata := float64(sum) / float64(len(arr))
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	var sumKuadrat float64
	for i := 0; i < len(arr); i++ {
		selisih := float64(arr[i]) - rata
		sumKuadrat += selisih * selisih
	}
	varians := sumKuadrat / float64(len(arr))
	stdDev := math.Sqrt(varians)
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	var cari int
	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cari)
	frekuensi := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("h. Frekuensi bilangan %d adalah: %d kali\n", cari, frekuensi)
}