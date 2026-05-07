package main
import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func nilaiPertama(arr arrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			return arr[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(arr arrayMahasiswa, n int, targetNIM string) int {
	max := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}
	return max
}

func main() {
	var n int
	var arraymahasiswa arrayMahasiswa
	var targetNIM string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arraymahasiswa[i].NIM, &arraymahasiswa[i].nama, &arraymahasiswa[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&targetNIM)

	pertama := nilaiPertama(arraymahasiswa, n, targetNIM)
	terbesar := nilaiTerbesar(arraymahasiswa, n, targetNIM)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", targetNIM, pertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", targetNIM, terbesar)
}