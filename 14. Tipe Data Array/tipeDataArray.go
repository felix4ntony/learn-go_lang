package main

import "fmt"

func main() {
	var names [3]string // wajib isi jumlah dari array tersebut
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"
	// names[3] = "Khannedy" //index 3 out of bounds

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		85,
		95,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(lagi[0])
	fmt.Println(len(lagi))
}
