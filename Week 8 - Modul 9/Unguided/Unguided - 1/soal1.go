package main
import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	radius int
}
func jarak(p,q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam (c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var lingkaran1, lingkaran2 lingkaran
	var p titik

	fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)
	fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)
	fmt.Scan(&p.x, &p.y)

	dilingkaran1 := didalam(lingkaran1, p)
	dilingkaran2 := didalam(lingkaran2, p)

	if dilingkaran1 && dilingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if dilingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	}else if dilingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	}else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}