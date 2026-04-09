package main 
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var maksimalHari int
	if tujuan == "domestik" {
		maksimalHari = 3
	} else if tujuan == "mancanegara" {
		maksimalHari = 8
	}

	if jumlahHari > maksimalHari {
		return maksimalHari
	}
	return jumlahHari
}
func biayaPerHari(jumlahMhs int) int {
	biayaSatuOrang := (2 * 35000) + 250000 + 300000
	return biayaSatuOrang * jumlahMhs
}
func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)
	biayaDomestikPerHari := biayaPerHari(jumlahMhs)
	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDomestikPerHari)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung * biayaDomestikPerHari) * 1.5
	}
}
func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour(domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah,lama,tujuan,&biaya)
	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U: Rp.%.0f",biaya)
}