package main
import "fmt"

func tanggunganHari(tujuan string) int {
	if tujuan == "domestik" {
		return 3
	} else if tujuan == "mancanegara" {
		return 8
	}
	return 0
}

func biayaPerHari() int {
	makan := 35000
	penginapan := 250000
	return makan + penginapan
}

func perhitunganBiaya(jumlahMahasiswa, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	uangSaku := 300000

	hari := tanggunganHari(tujuan)
	biayaHarian := biayaPerHari()

	total := jumlahMahasiswa * hari * biayaHarian
	total = total + (jumlahMahasiswa * uangSaku)

	*totalBiaya = float64(total)
}

func main() {
	var jumlahMahasiswa, lamaPerjalanan int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlahMahasiswa)
	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lamaPerjalanan)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	fmt.Println()

	fmt.Print("Biaya perjalanan yang harus dikeluarkan Tel-U: ")
	fmt.Scan(&biaya)
}