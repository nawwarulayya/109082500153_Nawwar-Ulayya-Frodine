package main
import "fmt"

func main() {
	var suara [21]int
	var x int

	totalMasuk := 0
	suaraSah := 0

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	ketua := 0
	wakil := 0
	max1 := -1
	max2 := -1

	for i := 1; i <= 20; i++ {

		if suara[i] > max1 {
			max2 = max1
			wakil = ketua

			max1 = suara[i]
			ketua = i

		} else if suara[i] > max2 {
			max2 = suara[i]
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}