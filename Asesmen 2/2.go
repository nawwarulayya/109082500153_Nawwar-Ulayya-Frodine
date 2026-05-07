package main
import "fmt"

const nMax int = 51
type mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}
type arrayMahasiswa [nMax]mahasiswa

func main() {
	var A arrayMahasiswa
	var N, idxFirst, idxMax int
	var nimCari string

	inputData(&A, &N)

	fmt.Print("masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	idxFirst = findFirst(A, N, nimCari)
	idxMax = findMax(A, N, nimCari)

	if idxFirst != -1 {
		fmt.Println("nilai pertama dari NIM", nimCari, "adalah", A[idxFirst].Nilai)
		fmt.Println("nilai terbesar dari NIM", nimCari, "adalah", A[idxMax].Nilai)
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}

func inputData(T *arrayMahasiswa, N *int) {
	fmt.Print("masukkan jumlah data: ")
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Print("masukkan data ke-", i+1, " : ")
		fmt.Scan(&T[i].NIM, &T[i].Nama, &T[i].Nilai)
	}
}

func findFirst(T arrayMahasiswa, N int, nim string) int {
	found := -1
	i := 0

	for i < N && found == -1 {
		if T[i].NIM == nim {
			found = i
		}
		i++
	}
	return found
}

func findMax(T arrayMahasiswa, N int, nim string) int {
	first := findFirst(T, N, nim)

	if first == -1 {
		return -1
	}
	idxMax := first

	for i := first; i < N; i++ {
		if T[i].NIM == nim && T[i].Nilai > T[idxMax].Nilai {
			idxMax = i
		}
	}
	return idxMax
}