package main

import "fmt"

func main() {
	const jumlahCalon = 20
	var suara [21]int 
	var totalMasuk, totalSah int
	var input int

	for {
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		totalMasuk++

		if input >= 1 && input <= 20 {
			suara[input]++
			totalSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	ketua := 0
	wakil := 0

	for i := 1; i <= jumlahCalon; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if suara[i] > suara[wakil] && i != ketua {
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}