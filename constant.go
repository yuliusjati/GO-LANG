package main

import "fmt"

func main(){
	const firstname string= "Yulius"
	const lastname = "Jati"
	const value = 1000

	fmt.Println(firstname)
	fmt.Println(lastname)
	// fmt.Println(value)

	const(
		namadepan = "Mpin"
		namabelakang = "Kepin"
	)
	fmt.Println(namadepan)
	fmt.Println(namabelakang)
}