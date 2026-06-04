package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	nama1  string
	nama2  string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func selectionSort(T *arrPemain, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if T[j].gol > T[idxMax].gol {
				idxMax = j
			} else if T[j].gol == T[idxMax].gol && T[j].assist > T[idxMax].assist {
				idxMax = j
			}
		}
		temp := T[i]
		T[i] = T[idxMax]
		T[idxMax] = temp
	}
}

func main() {
	var T arrPemain
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama1, &T[i].nama2, &T[i].gol, &T[i].assist)
	}

	selectionSort(&T, n)

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%v %v %v %v\n", T[i].nama1, T[i].nama2, T[i].gol, T[i].assist)
	}
}