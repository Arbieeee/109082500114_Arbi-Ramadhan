package main

import "fmt"

func tampilkanBarisan(n int) {
	if n == 1 {
		fmt.Printf("%d ", n)
		return
	}

	fmt.Printf("%d ", n)
	tampilkanBarisan(n - 1)
	fmt.Printf("%d ", n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	if n > 0 {
		tampilkanBarisan(n)
		fmt.Println()
	} else {
		fmt.Println("Masukkan bilangan bulat positif.")
	}
}