package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	f := func(x int) int {
		return x * x
	}
	g := func(x int) int {
		return x - 2
	}
	h := func(x int) int {
		return x + 1
	}
	fogoh := func(x int) int {
		return f(g(h(x)))
	}
	gohof := func(x int) int {
		return g(h(f(x)))
	}
	hofog := func(x int) int {
		return h(f(g(x)))
	}
	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}