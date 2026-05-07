package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) (x bool) {
	x = false
	for i := 0; i < n; i++ {
		if val == T[i] {
			x = true
		}
	}
	return x
}

func inputSet(T *set, n *int) {
	var val int
	*n = 0
	for {
		fmt.Scan(&val)
		if exist(*T, *n, val) {
			break
		}
		(*T)[*n] = val
		*n++
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if T1[i] == T2[j] {
				T3[*h] = T1[i]
				*h = *h + 1
				break
			}
		}
	}
}

func printSet (T set, n int){
	for i := 0; i < n; i++ {
			fmt.Print (T[i], " ")
	}
}

func main() {
	var T1, T2, T3 set
	var n, m, h int

	inputSet(&T1, &n)

	inputSet(&T2, &m)

	findIntersection(T1, T2, n, m, &T3, &h)
	printSet(T3, h)
}