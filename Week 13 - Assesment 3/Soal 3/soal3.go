package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai = [NMAX]partai

func posisi(t *tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func insertionSort(t *tabPartai, n int) {
	for i := 1; i < n; i++ {
		key := t[i]
		j := i - 1
		for j >= 0 && t[j].suara < key.suara {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func main() {
	var p tabPartai
	var input int
	n := 0

	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		idx := posisi(&p, n, input)
		if idx == -1 {
			p[n].nama = input
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
	}

	insertionSort(&p, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
	}
	if n > 0 {
		fmt.Println()
	}
}