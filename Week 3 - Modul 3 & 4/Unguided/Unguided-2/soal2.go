package main

import "fmt"

func tarifKendaraan(jenis string, jamMasuk int, durasi int) int {
	totalBiaya := 0
	jamSaatIni := jamMasuk

	for i := 1; i <= durasi; i++ {
		var tarif int
		if jenis == "motor" {
			if jamSaatIni >= 1 && jamSaatIni < 17 {
				tarif = 4000
			} else {
				tarif = 5000
			}
		} else if jenis == "mobil" {
			if jamSaatIni >= 1 && jamSaatIni < 17 {
				tarif = 6000
			} else {
				tarif = 7000
			}
		}

		totalBiaya = totalBiaya + tarif
		
		jamSaatIni++
		if jamSaatIni > 24 {
			jamSaatIni = 1
		}
	}
	return totalBiaya
}

func main() {
	var kendaraan string
	var masuk, keluar int
	var urutan int
	var durasi, total int
	urutan = 1

	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	for {
		fmt.Println("Kendaraan ", urutan)
		fmt.Print("Kendaraan apa? (motor/mobil/ - untuk selesai): ")
		fmt.Scan(&kendaraan)

		if kendaraan == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		if keluar >= masuk {
			durasi = keluar - masuk
		} else {
			durasi = (24 - masuk) + keluar
		}

		biaya := tarifKendaraan(kendaraan, masuk, durasi)

		fmt.Printf("Biaya parkir %s %d: %d\n", kendaraan, urutan, biaya)
		fmt.Println("========================================")
		fmt.Println()

		total = total + biaya

		urutan++
	}
	fmt.Printf("*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n",total)
}
