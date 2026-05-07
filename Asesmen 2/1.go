package main
import "fmt"

type set [2022]int
func exist(T set, x int, y int) bool {
	var i int = 0
	var status bool = false
	for i < x && !status {
		status = T[i] == y
		i++
	}
	return status
}

func inputSet(T *set, x *int){
	*x = 0
	var bilangan int
	fmt.Scan(&bilangan)
	for *x < 2022 && !exist(*T,*x,bilangan) {
		T[*x] = bilangan
		*x++
		fmt.Scan(&bilangan)
	}
}

func findInter(T1,T2 set, x,z int, T3 *set, h *int){
	var j int = 0
	*h = 0
	for j < x {
		if exist(T2,z, T1[j]) {
			T3[*h] = T1[j]
			*h++
		}
		j++
	}
}
func printSet(T set, x int) {
	var i int = 0
	for i < x {
		fmt.Print(T[i]," ")
		i++
	}
}

func main(){
	var s1,s2,s3 set
	var i1,i2,i3 int
	inputSet(&s1,&i1)
	inputSet(&s2,&i2)
	findInter(s1,s2,i1,i2,&s3,&i3)
	printSet(s3,i3)
}