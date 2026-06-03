package main
import "fmt"

func selectionSortAsc(arr []int, n int) {
	var idxMin, i, j int
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		if idxMin != i {
			arr[i], arr[idxMin] = arr[idxMin], arr[i]
		}
	}
}

func selectionSortDesc(arr []int, n int) {
	var idxMax, i, j int
	for i = 0; i < n-1; i++ {
		idxMax = i
		for j = i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		if idxMax != i {
			arr[i], arr[idxMax] = arr[idxMax], arr[i]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		ganjil := make([]int, m)
		genap := make([]int, m)
		cntGanjil := 0
		cntGenap := 0

		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil[cntGanjil] = x
				cntGanjil++
			} else {
				genap[cntGenap] = x
				cntGenap++
			}
		}

		selectionSortAsc(ganjil, cntGanjil)
		selectionSortDesc(genap, cntGenap)

		first := true
		for j := 0; j < cntGanjil; j++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			first = false
		}
		for j := 0; j < cntGenap; j++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			first = false
		}
		fmt.Println()
	}
}