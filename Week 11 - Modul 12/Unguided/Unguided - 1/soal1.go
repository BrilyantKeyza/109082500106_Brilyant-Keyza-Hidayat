package main
import "fmt"

func main() {
	var n int
	var totalMasuk, suaraSah int
	var vote [21] int

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
			fmt.Printf("%d: %d\n", i, vote[i])
		}
	}

}