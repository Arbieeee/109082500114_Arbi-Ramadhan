package main 

import "fmt"

const max = 100

type arr [max]int

func BinarySearch(a arr, n int, x int) bool{
	var kiri int = 0
	var kanan int = n - 1 
	var found bool = false
	var tengah int

	for kiri <= kanan && !found{
		tengah = (kiri + kanan) / 2
		if a[tengah] < x {
			kiri = tengah + 1
		} else if a[tengah] > x {
			kanan = tengah - 1
		} else {
			found = true
		}
	}
	return found
}

func main() {
	var data arr
	var n, x int

	fmt.Print("Input jumlah data: ")
	fmt.Scan(&n)

	fmt.Println("Input data ascending")
	for i := 0; i <= n; i++ {
		fmt.Scanln(&data[i])
	}

	fmt.Println("Data yang dicari: ")
	fmt.Scan(&x)

	if BinarySearch(data, n, x){
		fmt.Println("Data ditemukan")
	} else {
		fmt.Println("DATA TIDAK ditemukan")
	}
}