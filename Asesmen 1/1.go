package main
import "fmt"
const pi float64 = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, massaJenis float64) float64 {
	return volume(r, t) * massaJenis
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		fmt.Println("Selisih massa zat cair kiri dan zat cair kanan: ", m2 - m1)
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukkan jari-jari alas tabung: ")
	fmt.Scanln(&r)
	fmt.Print("Masukkan tinggi zat cair tabung kiri: ")
	fmt.Scanln(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri: ")
	fmt.Scanln(&mjKiri)
	fmt.Print("Masukkan tinggi zat cair tabung kanan: ")
	fmt.Scanln(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan: ")
	fmt.Scanln(&mjKanan)
	
	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)
	display(massaKanan, massaKiri)
}