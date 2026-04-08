package main
import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n - 1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan input: ")
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}