package main
import "fmt"

const nProv int = 10
type namaProv [nProv] string
type popProv [nProv] int
type tumbuhProv [nProv] float64

func InputData(prov *namaProv, pop *popProv, tumbuh *tumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d: ", i+1)
		fmt.Scan(&(*prov)[i], &(*pop)[i], &(*tumbuh)[i])
	} 
}

func provinsiTercepat(tumbuh tumbuhProv) int {
	tertinggi := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[tertinggi] {
			tertinggi = i
		}
	}
	return tertinggi
}

func prediksi(prov namaProv, pop popProv, tumbuh tumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan (Pertumbuhan > 2%) ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil := (tumbuh[i] + 1) * float64(pop[i])
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}	
}

func indexProvinsi(prov namaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}
func main() {
	var prov namaProv
	var pop popProv
	var tumbuh tumbuhProv
	var namaCari string
	
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&namaCari)

	tercepat := provinsiTercepat(tumbuh)
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat: ", prov[tercepat])

	mencari := indexProvinsi(prov, namaCari)
	if mencari != -1 {

		fmt.Println("Data provinsi yang dicari: ",prov[mencari])
	}else {
		fmt.Println("Data provinsi tidak ditemukan")
	}

	prediksi(prov,pop,tumbuh)

	


}