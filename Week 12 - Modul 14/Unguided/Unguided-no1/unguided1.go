package main

import "fmt"

func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&rumah[i])
		}

		selectionSort(rumah, m)

		for i := 0; i < m; i++ {
			fmt.Printf("%d ", rumah[i])
		}
		fmt.Println()
	}
}