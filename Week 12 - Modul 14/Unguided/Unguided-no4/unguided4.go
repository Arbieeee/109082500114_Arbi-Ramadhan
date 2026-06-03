package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

// DaftarkanBuku mengisi array pustaka dengan sejumlah n data buku dari masukan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n) // baca jumlah buku
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis,
			&pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// CetakTerfavorit menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	idxMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	fmt.Printf("%s %s %s %d\n", pustaka[idxMax].judul, pustaka[idxMax].penulis,
		pustaka[idxMax].penerbit, pustaka[idxMax].tahun)
}

// UrutBuku mengurutkan array pustaka berdasarkan rating secara menurun (descending)
// menggunakan algoritma insertion sort
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Cetak5Terbaru menampilkan 5 judul buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < batas {
		batas = n
	}
	for i := 0; i < batas; i++ {
		fmt.Printf("%s ", pustaka[i].judul)
	}
	fmt.Println()
}

// CariBuku mencari dan menampilkan buku dengan rating tertentu menggunakan binary search
func CariBuku(pustaka DaftarBuku, n int, r int) {
	// Binary search pada array yang sudah terurut descending berdasarkan rating
	kiri, kanan := 0, n-1
	ditemukan := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			// Tampilkan data buku yang ditemukan
			fmt.Printf("%s %s %s %d %d %d\n", pustaka[tengah].judul, pustaka[tengah].penulis,
				pustaka[tengah].penerbit, pustaka[tengah].tahun, pustaka[tengah].eksemplar, pustaka[tengah].rating)
			ditemukan = true
			break
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n int
	DaftarkanBuku(&pustaka, &n)

	// Cetak buku terfavorit (rating tertinggi) sebelum diurutkan
	CetakTerfavorit(pustaka, n)

	// Urutkan buku berdasarkan rating (menurun)
	UrutBuku(&pustaka, n)

	// Cetak 5 judul buku dengan rating tertinggi
	Cetak5Terbaru(pustaka, n)

	// Baca rating yang dicari
	var ratingDicari int
	fmt.Scan(&ratingDicari)

	// Cari buku dengan rating tersebut
	CariBuku(pustaka, n, ratingDicari)
}