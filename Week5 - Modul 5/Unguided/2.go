package main
import "fmt"

func polaBintang(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	polaBintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	polaBintang(n, 1)
}