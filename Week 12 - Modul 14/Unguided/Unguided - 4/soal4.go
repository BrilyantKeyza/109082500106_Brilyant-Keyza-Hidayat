package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	var i int
	for i = 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var idxMax, i int
	idxMax = 0
	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	fmt.Println(pustaka[idxMax].judul, pustaka[idxMax].penulis, pustaka[idxMax].penerbit, pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int
	for i = 1; i <= n-1; i++ {
		temp = pustaka[i]
		j = i
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i int
	batas := 5
	if n < 5 {
		batas = n
	}
	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var mid, low, high int
	low = 0
	high = n - 1
	mid = -1

	for low <= high {
		mid = (low + high) / 2
		if pustaka[mid].rating == r {
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if mid >= 0 && pustaka[mid].rating == r {
		fmt.Println(pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var r int

	DaftarkanBuku(&pustaka, &nPustaka)

	fmt.Scan(&r)

	CetakTerfavorit(pustaka, nPustaka)

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)

	CariBuku(pustaka, nPustaka, r)
}