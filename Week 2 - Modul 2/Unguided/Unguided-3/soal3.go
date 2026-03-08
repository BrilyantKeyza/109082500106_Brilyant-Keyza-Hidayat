package main
import "fmt"

func main() {
	var gram, kg int
	var sisaGram, biayaKg, biayaSisa, totalBiaya int 
	
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	sisaGram = gram % 1000
	biayaKg = kg * 10000

	if kg >= 10 {
		biayaSisa = 0
	}else if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	}else {
		biayaSisa = sisaGram * 15
	}
	totalBiaya = biayaKg + biayaSisa

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat: %d kg + %d gram", kg, sisaGram)
	fmt.Printf("\nDetail biaya: Rp. %d + Rp. %d ", biayaKg, biayaSisa)
	fmt.Print("\nTotal biaya: Rp. ",totalBiaya )
}