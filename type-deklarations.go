package main

import "fmt"

func main(){
	type nobuk string
	type book bool

	var buku1 nobuk = "0987"
	var kembali book = true

	fmt.Println(buku1)
	fmt.Println(kembali)
}