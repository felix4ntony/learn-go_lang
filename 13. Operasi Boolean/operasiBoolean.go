package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80
	fmt.Println(lulusAbsensi, lulusNilaiAkhir)

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}
