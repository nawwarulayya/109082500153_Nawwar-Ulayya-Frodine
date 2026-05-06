package main
import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah data harus antara 1 sampai 1000")
		return
	}

	var berat [1000]float64

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}