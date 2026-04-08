package main

import "fmt"

func barisan(n int) {
	fmt.Print(n, " ")

	if n == 1 {
		return
	}

	barisan(n - 1)

	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	barisan(n)
}