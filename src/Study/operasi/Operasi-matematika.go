package main

import "fmt"

func main(){
	var hasil = 20 + 60
	fmt.Println(hasil)

	var a = 35
	a /=5
	// +=
	// -=
	// %=
	// *=

	var b = 40
	var c = b-a
	fmt.Println(c)

	// Unary Operator //
	// ++   a = a + 1
	// --   a = a - 1
	//  -   negatif
	// +    potitif
	// !    boolean kebalikan

	c++
	fmt.Println(c)
}