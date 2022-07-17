package main

import "fmt"

func main(){
	// var person map[string]string =map[string]string{}
	person := map[string]string{
		"name" : "yulius",
		"address" : "Malang",
	}

	//cara menambah atau merubah
	person["title"] = "Programer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar Go-Lang"
	book["author"] = "yulius"
	book["kode"] = "098"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "kode")

	fmt.Println(book)
	fmt.Println(len(book))
}