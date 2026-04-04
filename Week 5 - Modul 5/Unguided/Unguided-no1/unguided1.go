package main

import "fmt"

func cetakBintang(n int) {
	if n <= 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(n - 1)
}

func buatPola(n int, barisSekarang int) {
	if barisSekarang > n {
		return
	}

	cetakBintang(barisSekarang)
	fmt.Println()
	buatPola(n, barisSekarang+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	buatPola(n, 1)
}