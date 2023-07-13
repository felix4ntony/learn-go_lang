package main

import "fmt"

func main() {
	// type 1 ----------------------------------------------------
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	// type 2 ----------------------------------------------------
	var friendName = "Eko Kurniawan"
	fmt.Println(friendName)

	friendName = "Eko Khannedy"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	var age_int8 int8 = 30
	fmt.Println(age_int8)
	// type 3 ----------------------------------------------------
	var age_int int
	age_int = 30
	fmt.Println(age_int)
	// type 4 ----------------------------------------------------
	country := "indonesian"
	fmt.Println(country)
	//type 5 ----------------------------------------------------
	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName, lastName)

}
