package main

import "fmt"

func main() {
	type noKTP string

	var ktpEko noKTP = "111111111"
	fmt.Println(ktpEko)
	fmt.Println(noKTP("222222222"))
}
