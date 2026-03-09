package main

import "fmt"

func main() {
	var a, b, c, d string
	berhasil := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d : ", i+1)
		fmt.Scan(&a, &b, &c, &d)

		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu" {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL : ", berhasil)
}