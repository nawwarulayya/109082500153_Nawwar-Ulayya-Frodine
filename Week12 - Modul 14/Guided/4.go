package main 
import "fmt"

type mahasiswa struct { 
	nama string
	nim string 
	ipk float64
}

type arrMhs [100]mahasiswa

func insertionSortStruct(T *arrMhs, n int) { 
	var temp mahasiswa
	var i, j int

	for i = 1; i < n; i++ { 
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].nama < temp.nama { 
			T[j + 1] = T[j]
			j--
		}
		T[j + 1] = temp
	}
}

func main() {
	var data arrMhs 
	var n, i int
	
	fmt.Print("Masukkan jumlah mahasiswa: ") 
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("\nData mahasiswa ke-", i + 1)
		
		fmt.Print("Nama: ") 
		fmt.Scan(&data[i].nama)

		fmt.Print("NIM: ") 
		fmt.Scan(&data[i].nim)
		
		fmt.Print("IPK: ") 
		fmt.Scan(&data[i].ipk)
	}

	fmt.Println("\nData sebelum sorting: ") 
	for i = 0; i < n; i++ {
		fmt.Println(data[i].nama, data[i].nim, data[i].ipk)
	}

	insertionSortStruct(&data, n)

	fmt.Println("\nData setelah sorting (Descending berdasarkan Nama): ")
	for i = 0; i < n; i++ {
		fmt.Println(data[i].nama, data[i].nim, data[i].ipk)
	}
}