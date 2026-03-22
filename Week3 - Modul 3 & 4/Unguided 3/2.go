package main
import ("fmt"; "strings")

func hitungBiaya(jenis string, masuk, keluar int) int {
	var tarifMalam, tarifSiang int

	if jenis == "motor" {
		tarifMalam = 4000
		tarifSiang = 5000
	} else if jenis == "mobil" {
		tarifMalam = 6000
		tarifSiang = 7000
	}

	total := 0

	for i := masuk; i < keluar; i++ {
		if i < 5 {
			total += tarifMalam
		} else {
			total += tarifSiang
		}
	}
	return total	
}

func main() {
	var jenis string
	var masuk, keluar int
	totalPendapatan := 0
	no := 1

	fmt.Println("==== Rekap Tarif Parkir Per Hari ====")

	for {
		fmt.Printf("\n*Kendaraan %d\n", no)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		jenis = strings.ToLower(jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan jam masuk kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan jam keluar kendaraan (0-24): ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)
		totalPendapatan += biaya

		fmt.Printf("Biaya parkir %s %d: %d\n", jenis, no, biaya)
		fmt.Println("====================================")
		no++
	}
	fmt.Printf("\n*** Total pendapatan hari ini adalah %d ***\n", totalPendapatan)
}