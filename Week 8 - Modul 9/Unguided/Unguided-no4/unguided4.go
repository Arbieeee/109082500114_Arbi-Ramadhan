package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	*n = 0

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	for _, ch := range line {
		if *n >= NMAX {
			break
		}
		if ch == '.' {
			break
		}
		if ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Input teks: ")
	isiArray(&tab, &m)

	var asli tabel
	for i := 0; i < m; i++ {
		asli[i] = tab[i]
	}

	balikanArray(&tab, m)

	fmt.Print("Teks terbalik: ")
	cetakArray(tab, m)

	fmt.Printf("Palindrom: %t\n", palindrome(asli, m))
}