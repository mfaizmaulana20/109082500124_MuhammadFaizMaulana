package main

import "fmt"

func main() {
	var beratGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := kg * 10000

	biayaTambahan := 0
	if kg < 10 {
		if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5
		} else {
			biayaTambahan = sisaGram * 15
		}
	} else {
		biayaTambahan = 0
	}

	totalBiaya := biayaKg + biayaTambahan
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}