package main
import "fmt"

const NMAX = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	for {
		var token string
		fmt.Scan(&token)
		if token == "." || *n >= NMAX {
			break
		}
		(*t)[*n] = rune(token[0])
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
	(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	asli := t
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if asli[i] != t[i] {
		return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Input teks : ")
	isiArray(&tab, &m)

	var tabBalik tabel
	for i := 0; i < m; i++ {
		tabBalik[i] = tab[i]
	}

	balikanArray(&tabBalik, m)

	fmt.Print("Teks terbalik : ")
	cetakArray(tabBalik, m)

	isPal := palindrom(tab, m)

	fmt.Print("Palindrom : ")
	if isPal {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}