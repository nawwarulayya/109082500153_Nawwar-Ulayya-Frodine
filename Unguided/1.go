package main
import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	kabisat := false

	if tahun%4 == 0 && tahun%100 != 0 || tahun%400 == 0 {
		kabisat = true
	}

	fmt.Println("Kabisat:", kabisat)
}