package main
import "fmt"

type array [1000]float64

func totalPerWadah(x, y int, beratIkan array) (float64, int) {
	var beratPerWadah float64
	var totalSemuaIkan float64
	var jumlahWadah int

	fmt.Println("====== Hasil ======")
	fmt.Print("Total Berat Ikan di Setiap Wadah : ")
	for i := 0; i < x; i++ {
		beratPerWadah += beratIkan[i]
		totalSemuaIkan += beratIkan[i]

		if (i+1)%y == 0 || i == x-1 {
			fmt.Printf("%.2f ", beratPerWadah)
			jumlahWadah++
			beratPerWadah = 0
		}
	}
	fmt.Println()
	return totalSemuaIkan, jumlahWadah
}

func hitungRatarata(total float64, jumlah int) float64 {
	if jumlah == 0 {
		return 0
	}
	return total / float64(jumlah)
}

func main() {
	var x, y int
	var beratIkan array

	fmt.Print("Masukkan banyak ikan yang akan dijual (x) dan Kapasitas ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}
	totalBerat, totalWadah := totalPerWadah(x, y, beratIkan)

	rerata := hitungRatarata(totalBerat, totalWadah)
	fmt.Printf("Berat Rata - rata ikan: %.2f\n", rerata)

}