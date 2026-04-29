package main

import (
	"fmt"
	"math"
)

func tampilkanSemua(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func tampilkanGanjil(arr []int, n int) {
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func tampilkanGenap(arr []int, n int) {
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func tampilkanKelipatanX(arr []int, n int, x int) {
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
}

func hapusElemen(arr []int, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
}

func hitungRataRata(arr []int, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n)
}

func hitungStandarDeviasi(arr []int, n int) float64 {
	rata := hitungRataRata(arr, n)
	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		jumlahKuadrat += selisih * selisih
	}
	return math.Sqrt(jumlahKuadrat / float64(n))
}

func hitungFrekuensi(arr []int, n int, bilangan int) int {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var n, x, idxHapus, bilCari int
	var data [1000]int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Print("a. Isi array: ")
	tampilkanSemua(data[:], n)

	fmt.Print("b. Indeks ganjil: ")
	tampilkanGanjil(data[:], n)

	fmt.Print("c. Indeks genap: ")
	tampilkanGenap(data[:], n)

	fmt.Print("d. Masukkan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	tampilkanKelipatanX(data[:], n, x)

	fmt.Printf("f. Rata-rata: %.2f\n", hitungRataRata(data[:], n))

	fmt.Printf("g. Standar Deviasi: %.2f\n", hitungStandarDeviasi(data[:], n))

	fmt.Print("h. Masukkan bilangan untuk cek frekuensi: ")
	fmt.Scan(&bilCari)
	fmt.Printf("   Frekuensi %d: %d\n", bilCari, hitungFrekuensi(data[:], n, bilCari))

	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	hapusElemen(data[:], &n, idxHapus)
	fmt.Print("   Array setelah dihapus: ")
	tampilkanSemua(data[:], n)
}