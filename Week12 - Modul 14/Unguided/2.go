package main 
import "fmt"

func selectionSortAsc(arr []int, n int) {
	for i := 0; i < n-1; i++ { 
		idxMin := i
		for j := i+1; j < n; j++ { 
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		temp := arr[idxMin] 
		arr[idxMin] = arr[i]
		arr[i] = temp
	}
}

func selectionSortDesc(arr []int, n int) { 
	for i := 0; i < n-1; i++ {
	idxMax := i
		for j := i+1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		temp := arr[idxMax] 
		arr[idxMax] = arr[i]
		arr[i] = temp
	}
}

func main() {
	var n int 
	fmt.Scan(&n)

	for i := 0; i < n; i++ { 
		var m int 
		fmt.Scan(&m)
	
		var ganjil []int 
		var genap []int

		for j := 0; j < m; j++ { 
			var nomor int 
			fmt.Scan(&nomor)
			if nomor%2 != 0 {
				ganjil = append(ganjil, nomor)
			} else {
				genap = append(genap, nomor)
			}
		}

		selectionSortAsc(ganjil, len(ganjil)) 
		selectionSortDesc(genap, len(genap))
		first := true
		for j := 0; j < len(ganjil); j++ { 
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j]) 
			first = false
		}

		for j := 0; j < len(genap); j++ { 
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			first = false
		}
		fmt.Println()
	}
}