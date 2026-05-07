package main
import "fmt"

const nProv int = 34
type NamaProv [nProv] string
type PopProv [nProv] int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv ){
	var x int
	for x = 0 ; x < nProv; x++{
		fmt.Scan(&prov[x],&pop[x],&tumbuh[x])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv ) int {var idx int = 0
	var x int
	for x = 1; x < nProv; x++ {
		if tumbuh[idx] < tumbuh[x] {
			idx = x
		}
	}
	return idx
}

func Prediksi (prov NamaProv, pop PopProv, tumbuh TumbuhProv){
	var x int
	var result float64
	for x = 0; x < nProv; x++{
		if tumbuh[x] > 0.02 {
			result = (tumbuh[x] + 1) * float64(pop[x])
			fmt.Println(prov[x], result)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	var found int = -1
	var x int = 0
	for x < nProv && found == -1 {
		if prov[x] == nama {
			found = x
		}
		x++
	}
	return found
}

func main(){
	var TProvinsi NamaProv
	var TPopulasi PopProv
	var TPertumbuhan TumbuhProv
	var cari string
	var idxTercepat, idxProvinsi int
	InputData(&TProvinsi, &TPopulasi, &TPertumbuhan)
	fmt.Scan(&cari)
	idxTercepat = ProvinsiTercepat(TPertumbuhan)
	fmt.Println(TProvinsi[idxTercepat])
	idxProvinsi = IndeksProvinsi(TProvinsi, cari)
	fmt.Println(TProvinsi[idxProvinsi])
	Prediksi(TProvinsi, TPopulasi, TPertumbuhan)
}