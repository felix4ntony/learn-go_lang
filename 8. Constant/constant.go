package main

import "fmt"

func main() {
	// type 1 ------------------------------------------------
	// const firstName = "Eko"
	// const lastName = "Khannedy"
	// const value = 1000
	// type 2 ------------------------------------------------
	const (
		firstName string = "Eko"
		lastName         = "Khannedy"
		value            = 1000
	)
	// -------------------------------------------------------

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
