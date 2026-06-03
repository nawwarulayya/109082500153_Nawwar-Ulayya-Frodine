package main 
import "fmt"

func selectionSort(arr []int, n int) { 
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

func main() {
	var n int 
	fmt.Scan(&n)

	for i := 0; i < n; i++ { 
		var m int 
		fmt.Scan(&m)
		
		var rumah []int
		for j := 0; j < m; j++ { 
			var nomor int 
			fmt.Scan(&nomor)
			rumah = append(rumah, nomor)
		
		}

		selectionSort(rumah, m)

		for j := 0; j < m; j++ { 
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}