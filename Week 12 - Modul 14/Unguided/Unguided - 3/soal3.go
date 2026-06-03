package main

import "fmt"

type arrInt [100]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	for i = 1; i <= n-1; i++ {
		temp = T[i]
		j = i
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
	}
}

func cekJarak(T *arrInt, n int) bool {
	var jarak, i int
	jarak = T[1] - T[0]
	for i = 2; i < n; i++ {
		if T[i]-T[i-1] != jarak {
			return false
		}
	}
	return true
}

func main() {
	var data arrInt
	var n, x int

	n = 0
	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data[n] = x
		n++
	}

	insertionSort(&data, n)

	var i int
	for i = 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	if cekJarak(&data, n) {
		jarak := data[1] - data[0]
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}