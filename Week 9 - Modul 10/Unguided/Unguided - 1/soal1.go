package main
import "fmt"

type array [1000] float64
func cariMin(n int, beratKelinci array)float64 {
	var min float64 = beratKelinci[0]
	var j int
	
	for j < n {
		if min > beratKelinci[j]{
			min = beratKelinci[j]
		}
		j = j + 1
	}
	return min
}

func cariMax(n int, beratKelinci array)float64 {
	var max float64 = beratKelinci[0]
	var j int

	for j < n {
		if max < beratKelinci[j]{
			max = beratKelinci[j]
		}
		j = j + 1
	}
	return max
}

func main() {
	var n int
	var beratKelinci array
	var min, max float64

	fmt.Print("Masukkan banyaknya anak kelinci yang akan ditimbang: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	min = cariMin(n, beratKelinci)
	max = cariMax(n, beratKelinci)

	fmt.Println("====== Hasil ======")
	fmt.Println("Berat kelinci terkecil:", min)
	fmt.Println("Berat kelinci terbesar:", max)

}