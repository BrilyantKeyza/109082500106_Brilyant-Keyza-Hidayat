package main
import "fmt"

func faktorial(n int) int{
	var hasil int
	hasil = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int{
	var hasil int
	hasil = faktorial(n) / faktorial(n-r)
	return hasil
}

func kombinasi(n, r int) int{
	var hasil int
	hasil = faktorial(n) / (faktorial(r) * faktorial(n-r))
	return hasil
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan 4 bilangan bulat: ")
	fmt.Scan(&a, &b, &c, &d)

	permutasi1 := permutasi(a, c)
	kombinasi1 := kombinasi(a, c)
	permutasi2 := permutasi(b, d)
	kombinasi2 := kombinasi(b, d)

	fmt.Println("Hasil: ")
	fmt.Println(permutasi1, kombinasi1)
	fmt.Println(permutasi2, kombinasi2)

}