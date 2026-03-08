package main
import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&n)

	if n % 400 == 0 {
		fmt.Print("Kabisat: True")
	}else if n % 4 == 0 && n % 100 != 0 {	
		fmt.Print("Kabisat: True")
	}else {
		fmt.Print("Kabisat: False")
	}
}