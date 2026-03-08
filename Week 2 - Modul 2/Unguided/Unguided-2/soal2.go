package main
import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var hasil bool = true
	var percobaan int

	for percobaan = 1; percobaan <= 5; percobaan++ {
		fmt.Printf("Percobaan %d: ", percobaan)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)
		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			hasil = false
		}
	}
	fmt.Print("BERHASIL: ", hasil)
}