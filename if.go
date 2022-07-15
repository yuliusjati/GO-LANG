package main

import "fmt"

func main(){
	var nama string = "Yulius"

	if nama == "fadil"{
		fmt.Println("halo fadil")
	} else if nama == "Yulius"{
		fmt.Println("Halo Yulius")
	} else {
		fmt.Println("Siapa Kamu?")
	}

	if length := len(nama); length > 5{
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("nama Sudah Benar")
	}
}