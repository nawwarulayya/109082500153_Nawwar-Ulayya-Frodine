package main
import "fmt"

func SequentialSearch(arrBuah [5]string, dataDicari string) int {
	var idx_found int = -1
	for i := 0; i < len(arrBuah); i++ { 
		if arrBuah[i] == dataDicari {
			idx_found = i
			break
		}
	}
	return idx_found
}

func main() {
	var arrBuah [5]string

	for i := 0; i < len(arrBuah); i++ {
		fmt.Printf("Masukkan data buah indeks ke-%d : ", i)
		fmt.Scan(&arrBuah[i])
	}

	var dataCari string
	fmt.Print("Masukkan data buah yang ingin dicari : ")
	fmt.Scan(&dataCari)

	index_data := SequentialSearch(arrBuah, dataCari)

	if index_data != -1 { 
		fmt.Printf("Data %s ditemukan pada indeks ke-%d!\n", dataCari, index_data)
	} else {
		fmt.Printf("Data %s tidak ditemukan!\n", dataCari)
	}
}