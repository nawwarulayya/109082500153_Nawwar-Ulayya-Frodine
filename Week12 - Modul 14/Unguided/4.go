package main 
import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string 
	eksemplar, tahun, rating	int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) { 
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, 
			&pustaka[i].penerbit, &pustaka[i].eksemplar, 
			&pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) { 
	if n == 0 {
		return
	}
	maxRating := pustaka[0].rating 
	idx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			idx = i
		}
	}
	fmt.Printf("%s %s %s %d\n", pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) { 
	for i := 1; i < n; i++ {
		temp := pustaka[i] 
		j := i
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) { 
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ { 
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	found := false
	idx := -1
	for kiri <= kanan && !found {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			found = true
			idx = tengah
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if found {
		fmt.Printf("%s %s %s %d %d %d\n", pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun, pustaka[idx].eksemplar, pustaka[idx].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, r int

	DaftarkanBuku(&pustaka, &n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&r)
	CariBuku(pustaka, n, r)
}