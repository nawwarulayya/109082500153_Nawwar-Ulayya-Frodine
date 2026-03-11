package main
import "fmt"

func main() {
	var satu string
	var dua string
	var tiga string
	var temp string

	fmt.Print("Masukan input: ")
	fmt.Scan(&satu)
	fmt.Print("Masukan input: ")
	fmt.Scan(&dua)
	fmt.Print("Masukan input: ")
	fmt.Scan(&tiga)

	fmt.Println("Output Awal: " + satu + " " + dua + " " + tiga)
	
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output Akhir: " + satu + " " + dua + " " + tiga)
}