package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address){
	address.Country = "Indonesia"
}
func main() {
	var Alamat = Address{
		City:"Subang", 
		Province:"Jawa Barat", 
		Country: "",
	}
	
	ChangeAddressToIndonesia(&Alamat)

	fmt.Println(Alamat)

	//var alamat *Address = &Alamat
	//ChangeAddressToIndonesia(alamat)
	//fmt.Println(alamat)
}