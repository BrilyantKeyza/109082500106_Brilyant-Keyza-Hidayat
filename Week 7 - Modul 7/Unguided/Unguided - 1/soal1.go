package main
import "fmt"

type suhu float64
func CelciusToReamure(suhu suhu) suhu{
	reamur := suhu * 4.0/5.0
	return reamur
}
func CelciusToFahrenheit(suhu suhu) suhu{
	fahrenheit := (suhu * (9.0/5.0)) + 32
	return fahrenheit
}
func CelciusToKelvin(suhu suhu) suhu{
	kelvin := suhu + 273.15
	return kelvin
}

func main() {
	var celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&celcius)

	fmt.Printf("\n%.0f celcius = %g reamure", celcius, CelciusToReamure(celcius))
	fmt.Printf("\n%.0f celcius = %g fahrenheit", celcius, CelciusToFahrenheit(celcius))
	fmt.Printf("\n%.0f celcius = %g kelvin", celcius, CelciusToKelvin(celcius))
}