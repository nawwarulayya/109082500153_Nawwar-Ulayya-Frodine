package main

import "fmt"

func hitungBiayaParkir(jenisKendaraan string, jamMasuk, jamKeluar int) int {
	var  biaya int
	var jamSiang, jamMalam int

	if jenisKendaraan == "motor" {

		if jamMasuk < 17 && jamKeluar > 17 {
			jamSiang = 17 - jamMasuk
			jamMalam = jamKeluar - 17
			biaya = (jamSiang * 4000) + (jamMalam * 5000)

		} else if jamKeluar <= 17 {
			biaya = (jamKeluar - jamMasuk) * 4000

		} else {
			biaya = (jamKeluar - jamMasuk) * 5000
		}

	} else if jenisKendaraan == "mobil" {

		if jamMasuk < 17 && jamKeluar > 17 {
			jamSiang = 17 - jamMasuk
			jamMalam = jamKeluar - 17
			biaya = (jamSiang * 6000) + (jamMalam * 7000)

		} else if jamKeluar <= 17 {
			biaya = (jamKeluar - jamMasuk) * 6000

		} else {
			biaya = (jamKeluar - jamMasuk) * 7000
		}
	}

	return biaya
}

func main() {

	var totalPendapatan int
	fmt.Println("==== Rekap Tarif Parkir cafe Per Hari ====")

	for i :=1; ; i++{
		var jenisKendaraan string
		var jamMasuk, jamKeluar, biaya int

		
		
		fmt.Println("\n*Kendaraan", i)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai) : ")
		fmt.Scan(&jenisKendaraan)
		
		if jenisKendaraan == "-" {
			break
		} 
		fmt.Print("Masukan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&jamMasuk)

		fmt.Print("Masukan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&jamKeluar)

		biaya = hitungBiayaParkir(jenisKendaraan, jamMasuk, jamKeluar)
		fmt.Printf("\nBiaya Parkir %s %d: %d\n", jenisKendaraan, i, biaya)
		fmt.Println("================================================\n")


		totalPendapatan += biaya
	}
	fmt.Printf("\n** Total Pendapatan Parkir Hari Ini Adalah %d **", totalPendapatan)


}