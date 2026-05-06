package main
import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan x (jumlah ikan) dan y (kapasitas wadah): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	var ikan [1000]float64

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y

	var totalPerWadah [1000]float64

	index := 0
	for i := 0; i < jumlahWadah; i++ {
		sum := 0.0
		count := 0

		for j := 0; j < y && index < x; j++ {
			sum += ikan[index]
			index++
			count++
		}

		totalPerWadah[i] = sum
		fmt.Print(sum, " ")
	}

	totalSemua := 0.0
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalPerWadah[i]
	}

	rataRata := totalSemua / float64(jumlahWadah)

	fmt.Println()
	fmt.Println("Rata-rata:", rataRata)
}