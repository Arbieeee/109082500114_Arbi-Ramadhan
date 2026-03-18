package main

import (
	"fmt"
)

func hitungBiaya(jenis string, masuk, keluar int) int {
	var tarifPagi, tarifMalam int

	switch jenis {
		case "motor":
		tarifPagi = 4000 
		tarifMalam = 5000
	case "mobil":
		tarifPagi = 6000
		tarifMalam = 7000 
	default:
		return 0
	}
	totalBiaya := 0
	jamSekarang := masuk

	for jamSekarang != keluar {
		if jamSekarang >= 0 && jamSekarang < 5 {
			totalBiaya += tarifPagi
		} else if jamSekarang >= 5 && jamSekarang < 17 { 
			totalBiaya += tarifPagi
		} else if jamSekarang >= 17 && jamSekarang < 24 {
			totalBiaya += tarifMalam
		}
		jamSekarang++
		if jamSekarang == 24 {
			jamSekarang = 0
		}
	}
	return totalBiaya
}
func main() {
	var jenis string
	var masuk, keluar int
	totalPendapatan := 0
	nomorKendaraan := 1

	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	for {
		fmt.Printf("\n*Kendaraan %d\n", nomorKendaraan)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}
		if jenis != "motor" && jenis != "mobil" {
			fmt.Println("Jenis kendaraan tidak valid. Masukkan motor atau mobil.")
			continue
		}
		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		if masuk < 0 || masuk > 24 || keluar < 0 || keluar > 24 {
			fmt.Println("Jam tidak valid. Gunakan format 0-24.")
			continue
		}
		biaya := hitungBiaya(jenis, masuk, keluar)
		totalPendapatan += biaya

		fmt.Printf("Biaya parkir %s %d: %d\n", jenis, nomorKendaraan, biaya)
		fmt.Println("================================")
		nomorKendaraan++
	}
	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", totalPendapatan)
}