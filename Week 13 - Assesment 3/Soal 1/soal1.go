package main

import "fmt"

const NMAX = 1000000 

type arrInt [NMAX]int 

func sorting(T *arrInt, n int){
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if T[j] < T[minIdx] {
                minIdx = j
            }
        }
        T[i], T[minIdx] = T[minIdx], T[i]
    }
}


func median(T arrInt, n int) float64 {
	
	if n == 0 {
		return 0
	}

	if n % 2 != 0 {
        return float64(T[n/2])
    }
    return float64(T[n/2-1]+T[n/2]) / 2.0

}

func main() {
	var data arrInt 
	var x int    
	var n int = 0

	fmt.Scan(&x)

	for x != -5313541 && n < NMAX {
		if x == 0 {
			sorting(&data, n)
			fmt.Println(median(data, n))
		} else {
			data[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}