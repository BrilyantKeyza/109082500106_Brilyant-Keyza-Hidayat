package main
import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64{
	vol := pi * r * r * t
	return vol
}
func massa(r, t, p float64) float64{
	massa := volume(r, t) * p
	return massa
}
func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Print("BALANCE")
	}else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = selisih * -1
		}
		fmt.Print("Selisih massa zat cair kiri dan massa zat cair kanan: ",selisih)
	}
}
func main() {
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukkan jari jari alas tabung: ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi zat cair tabung kiri: ")
	fmt.Scan(&tKiri)

	fmt.Print("Masukkan massa jenis zat cair tabung kiri: ")
	fmt.Scan(&mjKiri)

	fmt.Print("Masukkan tinggi zat cair tabung kanan: ")
	fmt.Scan(&tKanan)

	fmt.Print("Masukkan massa jenis zat cair tabung kanan: ")
	fmt.Scan(&mjKanan)


	massaKiri = massa(r, tKiri, mjKiri )
	massaKanan = massa(r, tKanan, mjKanan )

	display(massaKiri,massaKanan)
}