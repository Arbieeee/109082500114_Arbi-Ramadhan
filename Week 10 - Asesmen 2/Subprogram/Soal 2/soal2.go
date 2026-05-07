package main

import "fmt"

const NMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [NMax]mahasiswa

func cariNilaiPertama(arr arrayMahasiswa, n int, nimCari string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nimCari {
			return arr[i].nilai
		}
	}
	return -1 
}

func cariNilaiTerbesar(arr arrayMahasiswa, n int, nimCari string) int {
	nilaiMax := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == nimCari {
			if arr[i].nilai > nilaiMax {
				nilaiMax = arr[i].nilai
			}
		}
	}
	return nilaiMax
}

func main() {
	var data arrayMahasiswa
	var N int
	var nimCari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	if N > NMax {
		N = NMax
	}

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(data, N, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(data, N, nimCari)

	if nilaiPertama == -1 {
		fmt.Printf("NIM %s tidak ditemukan\n", nimCari)
	} else {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiTerbesar)
	}
}