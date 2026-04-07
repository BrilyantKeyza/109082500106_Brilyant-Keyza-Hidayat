package main
import "fmt"

func barisBilangan(n int){
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	barisBilangan(n - 1)
	fmt.Print(n, " ")
}
func main(){
	var n int
	fmt.Scan(&n)
	barisBilangan(n)
}