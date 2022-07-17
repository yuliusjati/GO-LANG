package main

import "fmt"

//interface kosong sama dengan jika kalian
type apapun interface{
	//jadi kalian bisa memasukan apapun mau itu string, boolean, int, dll
}

func Ups(i int) interface{} {
	if i == 1{
		return 1
	} else if i == 2{
		return false
	} else {
		return "kaowkaowka"
	}
}

func main() {
	var ada interface{} = Ups(2)
	kosong := Ups(4)
	fmt.Println(kosong, ada)
}