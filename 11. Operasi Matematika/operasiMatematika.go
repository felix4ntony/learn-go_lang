package main

import "fmt"

func main() {
	// type 1 ----------------------------------------------
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	//type 2 -----------------------------------------------
	var result = 10 + 10
	fmt.Println(result)

	//type 3 Augmented Assignment --------------------------
	var i = 10
	i += 10
	fmt.Println(i)

	//type 4 Unary Operator --------------------------------
	i++ // i = i+1
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative, positive)

}
