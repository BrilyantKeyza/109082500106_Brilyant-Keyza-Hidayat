package main

import "fmt"

func selectionSort(arr []int) {
	var idx_min, i, j int
	for i = 0; i < len(arr); i++ { 
		idx_min = i;
		for j = i + 1; j < len(arr); j++ { 
			if arr[j] < arr[idx_min] {
				idx_min = j
			}
		}
		arr[i], arr[idx_min] = arr[idx_min], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}