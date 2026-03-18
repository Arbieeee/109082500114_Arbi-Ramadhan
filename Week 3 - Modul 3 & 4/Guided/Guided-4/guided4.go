package main

import (
	"fmt"
)

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
	fmt.Scan(&n)
	cetakDeret(n)
}