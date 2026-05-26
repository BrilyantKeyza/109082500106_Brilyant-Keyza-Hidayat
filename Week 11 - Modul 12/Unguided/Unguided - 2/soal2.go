package main
import "fmt"

func main() {
	var n int
	var totalMasuk, suaraSah int
	var vote [21]int
	var ketuaRT, wakil int

	for {
		fmt.Scan(&n)
 
		if n == 0 {
			break
		}
 
		totalMasuk++
 
		if n >= 1 && n <= 20 {
			suaraSah++
			vote[n]++
		}
	}
	fmt.Println("Suara Masuk: ", totalMasuk)
	fmt.Println("Suara Sah: ", suaraSah)

	for i := 1; i <= 20; i++ {
		if vote[i] > 0 {
			if ketuaRT == -1 || vote[i] > vote[ketuaRT] {
				ketuaRT = i
			}
		}
	}

	for i := 1; i <= 20; i++ {
		if vote[i] > 0 && i != ketuaRT {
			if wakil == -1 || vote[i] > vote[wakil] {
				wakil = i
			}
		}
	}
	fmt.Printf("Ketua RT: %d\n", ketuaRT)
	fmt.Printf("Wakil Ketua: %d\n", wakil)
}