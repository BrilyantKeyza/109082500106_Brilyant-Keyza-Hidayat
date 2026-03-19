package main
import "fmt"

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Println("Luas Persegi: ", luas)
	fmt.Println("Keliling Persegi: ", keliling)
}
func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Println("Luas Persegi: ", luas)
	fmt.Println("Keliling Persegi: ", keliling)
}
func hitungLingkaran(jarijari float64) {
	phi := 3.14
	luas := phi * (jarijari * jarijari)
	keliling := 2 * phi * jarijari
	fmt.Println("Luas Lingkaran: ", luas)
	fmt.Println("Keliling keliling: ", keliling)
}

func main() {
	var pilihan int
	var sisi, panjang, lebar int
	var jarijari float64

	fmt.Println("=== PROGRAM BANGUN DATAR ===")
	fmt.Println("1. Hitung Luas & Keliling Persegi")
	fmt.Println("2. Hitung Luas & Keliling Persegi Panjang")
	fmt.Println("3. Hitung Luas & Keliling Lingkaran")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Sisi: ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		fmt.Print("Masukkan Panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan Lebar: ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		fmt.Print("Masukkan Jari - Jari: ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Print("Pilihan tidak ada")
	}
}
