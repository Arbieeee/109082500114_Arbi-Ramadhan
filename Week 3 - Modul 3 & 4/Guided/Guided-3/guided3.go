package main

import (
	"fmt"
)

func hitungSkor(soal, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0
	
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 { 
			*soal++
			*skor += waktu
		}
	}
}
func main() {
	var nama, pemenang string
	var soal, skor int
	var maxSoal, minSkor int
	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		hitungSkor(&soal, &skor)

		if soal > maxSoal {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		} else if soal == maxSoal {
			if skor < minSkor {
				minSkor = skor
				pemenang = nama
			}
		}
	}
	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
}