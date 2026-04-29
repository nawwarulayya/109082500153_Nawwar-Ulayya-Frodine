package main

import (
	"fmt"
)

const NMAX = 127

type tabel [NMAX]rune

// isi array
func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scan(&ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

// cetak array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

// balik array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// cek palindrome
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Hves : ")
	isiArray(&tab, &n)

	// tampilkan array awal
	cetakArray(tab, n)

	// balik array
	balikanArray(&tab, n)

	fmt.Print("Rvuvs tves : ")
	cetakArray(tab, n)

	// cek palindrome
	if palindrome(tab, n) {
		fmt.Println("palindrom")
	} else {
		fmt.Println("bukan palindrom")
	}
}