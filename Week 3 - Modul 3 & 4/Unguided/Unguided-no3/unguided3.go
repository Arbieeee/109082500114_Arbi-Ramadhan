package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Printf("%d\n", n)
}
func main() {
	var n int

	for i := 0; i < 3; i++ {
		fmt.Print("Masukkan bilangan : ")
		fmt.Scan(&n)

		if n < 1 || n >= 1000000 {
			fmt.Println("Bilangan harus positif dan kurang dari 1000000")
			i-- 
			continue
		}
		cetakDeret(n)
	}
}