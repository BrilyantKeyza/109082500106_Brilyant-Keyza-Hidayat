package main

import "fmt"

func main() {
	var angka, pilihan int

	fmt.Println("======== MENU ========")
	fmt.Println("|1. Cek angka        |")
	fmt.Println("|2. Cek Ganjil Genap |")
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Angka: ")
		fmt.Scan(&angka)
		if angka == 10 {
			fmt.Print("Angka adalah 10")
		} else {
			fmt.Print("Angka bukan 10")
		}
	case 2:
		fmt.Print("Masukkan Angka: ")
		fmt.Scan(&angka)
		if angka%2 == 0 {
			fmt.Print("Angka Genap")
		} else if angka%2 != 0 {
			fmt.Print("Angka Ganjil")
		}

	default:
		fmt.Print("Pilihan Tidak Valid")
	}

}
