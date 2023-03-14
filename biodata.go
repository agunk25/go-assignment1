package main

import (
	"fmt"
	"os"
	"strings"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var args = os.Args
	d_name := []string{"Thomas", "Abe", "Rudy", "Malvin", "Agung"}
	d_alamat := []string{"Kota A", "Kota B", "Kota C", "Kota D", "Kota E"}
	d_pekerjaan := []string{"Programer", "UI/UX", "UI/UX", "Backend", "Backend"}

	var biodatas []Biodata
	for i, d := range d_name {
		var bio Biodata
		bio.Nama = d
		bio.Alamat = d_alamat[i]
		bio.Pekerjaan = d_pekerjaan[i]
		bio.Alasan = fmt.Sprint("Alasan ", d)

		biodatas = append(biodatas, bio)
	}

	getData(args[1], biodatas)
}

func getData(n string, b []Biodata) {
	for i, d := range b {
		if strings.ToLower(d.Nama) == strings.ToLower(n) {
			fmt.Println("ID : ", i)
			fmt.Println("nama : ", d.Nama)
			fmt.Println("alamat : ", d.Alamat)
			fmt.Println("pekerjaan : ", d.Pekerjaan)
			fmt.Println("alasan : ", d.Alasan)
		}
	}
}
