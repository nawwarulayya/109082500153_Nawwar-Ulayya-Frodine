package main
import "fmt"

type angka int
type kata string

type Buku struct {
	judul          kata
	penulis        kata
	penerbit       kata
	tahunTerbit    angka
	jumlahHalaman  angka
}

func main() {
	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	fmt.Scanln(&b.judul)

	fmt.Print("Masukkan nama penulis : ")
	fmt.Scanln(&b.penulis)

	fmt.Print("Masukkan penerbit : ")
	fmt.Scanln(&b.penerbit)

	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scanln(&b.tahunTerbit)

	fmt.Print("Masukkan jumlah halaman : ")
	fmt.Scanln(&b.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", b.judul)
	fmt.Println("Penulis :", b.penulis)
	fmt.Println("Penerbit :", b.penerbit)
	fmt.Println("Tahun Terbit :", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}