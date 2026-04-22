package main
import "fmt"

type angka int
type kata string

type Buku struct {
	judul kata
	penulis kata
	penerbit kata
	tahunTerbit angka
	jumlahHalaman angka
}
func main() {
	var buku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan Judul Buku: ")
	fmt.Scan(&buku.judul)
	fmt.Print("Masukkan Nama Penulis: ")
	fmt.Scan(&buku.penulis)
	fmt.Print("Masukkan Penerbit: ")
	fmt.Scan(&buku.penerbit)
	fmt.Print("Masukkan Tahun Terbit: ")
	fmt.Scan(&buku.tahunTerbit)
	fmt.Print("Masukkan Jumlah Halaman: ")
	fmt.Scan(&buku.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku:",buku.judul)
	fmt.Println("Penulis:",buku.penulis)
	fmt.Println("Penerbit:",buku.penerbit)
	fmt.Println("Judul Buku:",buku.tahunTerbit)
	fmt.Println("Judul Buku:",buku.jumlahHalaman)
}