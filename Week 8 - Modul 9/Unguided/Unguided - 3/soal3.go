package main
import "fmt"

func main() {
	var skorA, skorB int
	var klubA, klubB string
	var pemenang [100] string
	var jumlahData int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	PertandinganKe := 1
	for {
		fmt.Printf("Pertandingan %d: ", PertandinganKe)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[jumlahData] = klubA
		} else if skorB > skorA {
			pemenang[jumlahData] = klubB
		} else {
			pemenang[jumlahData] = "Draw"
		}
		PertandinganKe++
		jumlahData++
	}

	fmt.Println()
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}