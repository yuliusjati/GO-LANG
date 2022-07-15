package main

import "fmt"

func main() {
	name := "sutris "

	switch name {
	case "yulius":
		fmt.Println("hello Yulius")
	case "jati" :
		fmt.Println("hello Jati")
	default:
		fmt.Println("Hello")
	}

	//switch length := len(name); length > 5{
	//case true :
	//	fmt.Println("Nama terlalu panjang")
	//case false : 
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch{
	case length > 10: 
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama sudah cocok")
	}
}