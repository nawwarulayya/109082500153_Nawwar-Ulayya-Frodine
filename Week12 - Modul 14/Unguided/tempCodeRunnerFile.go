package main 
import "fmt"

func insertionSort(arr []int, n int) { 
	for i := 1; i < n; i++ {
		temp := arr[i] j := i
		for j > 0 && temp < arr[j-1] { 
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
	}
}

func main() {
	var arr []int 
	var input int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		arr = append(arr, input)
	}

	n := len(arr) 
	if n > 0 {
		insertionSort(arr, n)
	}
	
	for i := 0; i < n; i++ { 
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()

	if n < 2 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		jarak := arr[1] - arr[0]
		tetap := true
		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != jarak {
				tetap = false
				break
			}
		}
		if tetap {
			fmt.Printf("Data berjarak %d\n", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}