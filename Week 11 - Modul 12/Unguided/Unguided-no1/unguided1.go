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

	for i := 1; i <= jumlahCalon; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}