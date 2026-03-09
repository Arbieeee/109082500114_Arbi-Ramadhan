package main

import "fmt"

func main() {
	var beratGram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}
	if kg > 10 {
			biayaSisa = sisaGram * 5
			biayaSisa = 0
	}
	totalBiaya := biayaKg + biayaSisa
	fmt.Println("\n==== DEtail Perhitungan ====")
	fmt.Printf("Detail Berat : %d kg + %d gram\n", kg, sisaGram)
	fmt.Printf("Detain Biaya :Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total Biaya : Rp. %d\n", totalBiaya)
}