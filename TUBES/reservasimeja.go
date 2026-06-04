package main

import "fmt"


type Meja struct {
	Nomor         int
	Kapasitas     int
	Status        string 
	JumlahDipesan int    
}

type Pelanggan struct {
	ID   int
	Nama string
	NoHP string
}

type Reservasi struct {
	IDReservasi int
	IDPelanggan int
	NomorMeja   int
	Waktu       string 
}

func TambahMeja(daftarMeja *[]Meja, nomor int, kapasitas int) {
	mejaBaru := Meja{
		Nomor:         nomor,
		Kapasitas:     kapasitas,
		Status:        "Tersedia",
		JumlahDipesan: 0,
	}
	*daftarMeja = append(*daftarMeja, mejaBaru)
	fmt.Printf("Sukses: Meja %d berhasil ditambahkan.\n", nomor)
}

func UbahMeja(daftarMeja []Meja, nomor int, kapasitasBaru int) {
	for i := range daftarMeja {
		if daftarMeja[i].Nomor == nomor {
			daftarMeja[i].Kapasitas = kapasitasBaru
			fmt.Printf("Sukses: Kapasitas meja %d diubah menjadi %d.\n", nomor, kapasitasBaru)
			return
		}
	}
	fmt.Printf("Gagal: Meja %d tidak ditemukan.\n", nomor)
}

func HapusMeja(daftarMeja *[]Meja, nomor int) {
	for i, meja := range *daftarMeja {
		if meja.Nomor == nomor {
			*daftarMeja = append((*daftarMeja)[:i], (*daftarMeja)[i+1:]...)
			fmt.Printf("Sukses: Meja %d berhasil dihapus.\n", nomor)
			return
		}
	}
	fmt.Printf("Gagal: Meja %d tidak ditemukan.\n", nomor)
}

func TambahPelanggan(daftarPelanggan *[]Pelanggan, id int, nama string, nohp string) {
	plgBaru := Pelanggan{ID: id, Nama: nama, NoHP: nohp}
	*daftarPelanggan = append(*daftarPelanggan, plgBaru)
	fmt.Printf("Sukses: Pelanggan %s berhasil ditambahkan.\n", nama)
}

func BuatReservasi(daftarMeja []Meja, daftarReservasi *[]Reservasi, idReservasi int, idPelanggan int, nomorMeja int, waktu string) {
	for i := range daftarMeja {
		if daftarMeja[i].Nomor == nomorMeja {
			if daftarMeja[i].Status == "Tersedia" {
				daftarMeja[i].Status = "Dipesan"
				daftarMeja[i].JumlahDipesan++

				reservasiBaru := Reservasi{
					IDReservasi: idReservasi,
					IDPelanggan: idPelanggan,
					NomorMeja:   nomorMeja,
					Waktu:       waktu,
				}
				*daftarReservasi = append(*daftarReservasi, reservasiBaru)
				
				fmt.Printf("Sukses: Meja %d berhasil dipesan untuk waktu/tanggal %s.\n", nomorMeja, waktu)
				return
			} else {
				fmt.Printf("Gagal: Meja %d saat ini sedang %s.\n", nomorMeja, daftarMeja[i].Status)
				return
			}
		}
	}
	fmt.Printf("Gagal: Meja %d tidak ditemukan.\n", nomorMeja)
}

func sequentialSearchNomor(data []Meja, target int) int {
	for i := 0; i < len(data); i++ {
		if data[i].Nomor == target {
			return i
		}
	}
	return -1
}

func binarySearchNomor(data []Meja, target int) int {
	left := 0
	right := len(data) - 1

	for left <= right {
		mid := (left + right) / 2

		if data[mid].Nomor == target {
			return mid
		} else if data[mid].Nomor < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func selectionSortNomor(data []Meja) {
    n := len(data)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if data[j].Nomor < data[minIdx].Nomor {
                minIdx = j
            }
        }
        data[i], data[minIdx] = data[minIdx], data[i]
    }
}

func insertionSortKapasitas(data []Meja) {
	n := len(data)
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j].Kapasitas > temp.Kapasitas {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func statistikPerHari(data []Reservasi) {
	fmt.Println("\n=== Statistik Reservasi Per Hari ===")
	if len(data) == 0 {
		fmt.Println("Belum ada data reservasi.")
		return
	}

	statistik := make(map[string]int)

	for _, r := range data {
		statistik[r.Waktu]++
	}

	for tanggal, jumlah := range statistik {
		fmt.Printf("Waktu/Tanggal %s : %d reservasi\n", tanggal, jumlah)
	}
}

func mejaTerfavorit(data []Reservasi) {
	fmt.Println("\n=== Meja Paling Sering Dipesan ===")
	if len(data) == 0 {
		fmt.Println("Belum ada data reservasi.")
		return
	}

	frekuensiMeja := make(map[int]int)

	for _, r := range data {
		frekuensiMeja[r.NomorMeja]++
	}

	max := 0
	mejaFavorit := 0

	for meja, jumlah := range frekuensiMeja {
		if jumlah > max {
			max = jumlah
			mejaFavorit = meja
		}
	}

	fmt.Printf("Meja nomor %d paling sering dipesan sebanyak %d kali\n", mejaFavorit, max)
}

func main() {
	listMeja := []Meja{
		{Nomor: 1, Kapasitas: 2, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 2, Kapasitas: 4, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 3, Kapasitas: 6, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 4, Kapasitas: 2, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 5, Kapasitas: 8, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 6, Kapasitas: 4, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 7, Kapasitas: 10, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 8, Kapasitas: 2, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 9, Kapasitas: 6, Status: "Tersedia", JumlahDipesan: 0},
		{Nomor: 10, Kapasitas: 4, Status: "Tersedia", JumlahDipesan: 0},
	}
	
	var listPelanggan []Pelanggan
	var listReservasi []Reservasi
	var pilihan int

	for {
		fmt.Println("\n==================================")
		fmt.Println("   SISTEM RESERVASI RESTORAN")
		fmt.Println("==================================")
		fmt.Println("1. Tambah Meja")
		fmt.Println("2. Hapus Meja")
		fmt.Println("3. Tambah Pelanggan")
		fmt.Println("4. Buat Reservasi")
		fmt.Println("5. Cari Meja (Sequential Search)")
		fmt.Println("6. Urutkan Meja Berdasarkan Kapasitas")
		fmt.Println("7. Tampilkan Semua Meja")
		fmt.Println("8. Tampilkan Statistik (Poin E)")
		fmt.Println("9. Keluar Program")
		fmt.Println("==================================")
		fmt.Print("Pilih menu (1-9): ")
		
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var nomor, kapasitas int
			fmt.Print("Masukkan Nomor Meja: ")
			fmt.Scan(&nomor)
			fmt.Print("Masukkan Kapasitas Meja: ")
			fmt.Scan(&kapasitas)
			TambahMeja(&listMeja, nomor, kapasitas)

		case 2:
			var nomor int
			fmt.Print("Masukkan Nomor Meja yang ingin dihapus: ")
			fmt.Scan(&nomor)
			HapusMeja(&listMeja, nomor)

		case 3:
			var id int
			var nama, nohp string
			fmt.Print("Masukkan ID/NIK Pelanggan: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan Nama Pelanggan (tanpa spasi): ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan No HP Pelanggan: ")
			fmt.Scan(&nohp)
			TambahPelanggan(&listPelanggan, id, nama, nohp)

		case 4:
			var idRes, idPel, noMeja int
			var waktu string
			fmt.Print("Masukkan ID Reservasi (bebas, misal 1): ")
			fmt.Scan(&idRes)
			fmt.Print("Masukkan ID Pelanggan yang memesan: ")
			fmt.Scan(&idPel)
			fmt.Print("Masukkan Nomor Meja yang dipesan: ")
			fmt.Scan(&noMeja)
			fmt.Print("Masukkan Tanggal/Waktu (contoh: 12-Mei): ")
			fmt.Scan(&waktu)
			BuatReservasi(listMeja, &listReservasi, idRes, idPel, noMeja, waktu)

		case 5:
			fmt.Println("Metode pencarian:")
			fmt.Println("A. Sequential Search (Bisa untuk data tidak terurut)")
			fmt.Println("B. Binary Search (Pastikan data terurut berdasarkan Nomor Meja)")
			fmt.Print("Pilih (A/B): ")
			var metodeCari string
			fmt.Scan(&metodeCari)

			var target int
			fmt.Print("Masukkan Nomor Meja yang dicari: ")
			fmt.Scan(&target)

			var idx int
			if metodeCari == "A" || metodeCari == "a" {
				idx = sequentialSearchNomor(listMeja, target)
				fmt.Println(">> Menggunakan algoritma Sequential Search...")
			} else if metodeCari == "B" || metodeCari == "b" {
				idx = binarySearchNomor(listMeja, target)
				fmt.Println(">> Menggunakan algoritma Binary Search...")
			} else {
				fmt.Println("Pilihan tidak valid.")
				continue 
			}

			if idx != -1 {
				fmt.Printf("Meja %d ditemukan! Kapasitas: %d, Status: %s\n",
					listMeja[idx].Nomor, listMeja[idx].Kapasitas, listMeja[idx].Status)
			} else {
				fmt.Println("Meja tidak ditemukan.")
			}

		case 6:
			fmt.Println("Metode pengurutan:")
    		fmt.Println("A. Selection Sort (Berdasarkan Nomor Meja)") 
    		fmt.Println("B. Insertion Sort")
    		fmt.Print("Pilih (A/B): ")
    		var metode string
    		fmt.Scan(&metode)

    		if metode == "A" || metode == "a" {
    		    selectionSortNomor(listMeja) 
    		    fmt.Println("Berhasil diurutkan berdasarkan Nomor Meja dengan Selection Sort.")
    		} else if metode == "B" || metode == "b" {
    		    insertionSortKapasitas(listMeja)
    		    fmt.Println("Berhasil diurutkan dengan Insertion Sort.")
    		} else {
    		    fmt.Println("Pilihan tidak valid.")
    		}
			
		case 7:
			fmt.Println("\n--- DAFTAR MEJA ---")
			if len(listMeja) == 0 {
				fmt.Println("Belum ada data meja.")
			} else {
				for _, meja := range listMeja {
					fmt.Printf("Meja %d | Kapasitas: %d | Status: %s\n", meja.Nomor, meja.Kapasitas, meja.Status)
				}
			}

		case 8:
			statistikPerHari(listReservasi)
			mejaTerfavorit(listReservasi)

		case 9:
			fmt.Println("Terima kasih telah menggunakan program ReservaResto!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
