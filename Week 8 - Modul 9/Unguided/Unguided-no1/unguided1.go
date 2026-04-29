package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) < float64(c.r)
}

func main() {
	var cx1, cy1, r1 int
	fmt.Scan(&cx1, &cy1, &r1)
	ling1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, r: r1}

	var cx2, cy2, r2 int
	fmt.Scan(&cx2, &cy2, &r2)
	ling2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, r: r2}

	var x, y int
	fmt.Scan(&x, &y)
	titik := Titik{x: x, y: y}

	dalam1 := didalam(ling1, titik)
	dalam2 := didalam(ling2, titik)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}