package main
import "fmt"

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println("Luas Persegi: ", luas)
	fmt.Println("Keliling Persegi: ", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("Luas Persegi Panjang: ", luas)
	fmt.Println("Keliling Persegi Panjang: ", keliling)
}	

func hitungLingkaran(jari2 float64) {
	const phi = 3.14
	luas := phi * jari2 * jari2
	keliling := 2 * phi * jari2

	fmt.Printf("Luas Lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling Lingkaran: %.2f\n", keliling)
}

func main() {
	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("\nMasukkan sisi: ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)

	case 2:
		var panjang, lebar int
		fmt.Print("\nMasukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)

	case 3:
		var jari2 float64
		fmt.Print("\nMasukkan jari-jari: ")
		fmt.Scan(&jari2)
		hitungLingkaran(jari2)
	
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
