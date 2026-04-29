package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Array setelah dihapus:", arr)

	total := 0
	for _, v := range arr {
		total += v
	}
	mean := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	var sum float64
	for _, v := range arr {
		sum += math.Pow(float64(v)-mean, 2)
	}
	std := math.Sqrt(sum / float64(len(arr)))
	fmt.Println("Standar deviasi:", std)

	var cari, count int
	fmt.Print("Masukkan angka yang dicari: ")
	fmt.Scan(&cari)

	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}