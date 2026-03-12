# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[MUHAMMADFAIZMAULANA] - [109082500124]</p>

## Unguided 

### 2a. [Soal]
#### soal1.go

```go

```
### Output Unguided :
package main
import "fmt"
func main() {
var (
	satu, dua, tiga string
	temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)

fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)

fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)

fmt.Println("Output awal = ", satu, dua, tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = ", satu, dua, tiga)
}

##### Output 
![Screenshot Output Unguided 2_a](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana/blob/main/modul1/output/output-soal2a.png)
[penjelasan]:Jadi program ini tuh berguna untuk merubah posisi dari setiap variabelnya dimana temp
menyimpan nomor satu,satu menyimpan variabel nomor dua,dua menyimpan variabel
nomor tiga,dan tiga menyimpan variabel punya temp dimana si tempnya itu udah diganti
isinya dengan variabel nomor satu.

 2b. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	
	berhasil := true

	for i := 1; i <= 5; i++ {
		var g1, g2, g3, g4 string

		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&g1, &g2, &g3, &g4)

		if g1 != "merah" || g2 != "kuning" || g3 != "hijau" || g4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Printf("BERHASIL: %v\n", berhasil)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_b](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana/blob/main/modul1/output/output-soal2b.png)
[penjelasan]:Program ini menggunakan perulangan untuk menerima input empat warna sebanyak lima kali,
lalu membandingkan setiap set input tersebut dengan urutan standar "merah", "kuning",
"hijau", dan "ungu". Variabel penanda (berhasil) diinisialisasi sebagai true, namun akan diubah
menjadi false jika ditemukan satu saja percobaan yang urutan warnanya tidak sesuai, sehingga
program akan menghasilkan nilai true hanya jika kelima percobaan secara konsisten memiliki
urutan yang tepat.

 2c. [Soal]
#### soal1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_c](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana/blob/main/modul1/output/output-soal2c.png)
[penjelasan]:Program ini menghitung biaya pengiriman parsel dengan membagi total berat gram
menjadi jumlah kilogram dan sisa gram, di mana biaya dasar ditetapkan sebesar Rp10.000
per kilogram. Program kemudian menerapkan aturan khusus: jika berat total kurang dari 10
kg, sisa gram akan dikenakan biaya Rp5 per gram jika jumlahnya 500 gram atau lebih, dan
Rp15 per gram jika kurang dari 500 gram; namun, jika berat total sudah melebihi 10 kg,
maka sisa gram tersebut digratiskan, sehingga total biaya hanya dihitung berdasarkan berat
kilogramnya saja.