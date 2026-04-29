package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var pemenang []string
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Printf("Hasil %d : %s\n", pertandingan, klubA)
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			fmt.Printf("Hasil %d : %s\n", pertandingan, klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Printf("Hasil %d : Draw\n", pertandingan)
		}

		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
}