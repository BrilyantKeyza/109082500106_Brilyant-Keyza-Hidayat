package main
import "fmt"

func cetak(n int){
	if n <= 0 {
		return 
	}
	cetak(n - 1)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main(){
	var n int
	fmt.Scan(&n)
	cetak(n)
}