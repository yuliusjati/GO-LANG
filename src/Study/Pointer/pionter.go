package main

import "fmt"

type Address struct {
	City, Province, Country string
}
type Animal struct {
	Name, Kota, Kelamin string
}

func main() {
	//pass by value
	address1 := Address{"Malang", "jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Banyuwangi"

	fmt.Println(address1)//data tidak berubah
	fmt.Println(address2)

	//pass by reference
	var animal1 Animal = Animal{"buaya", "jakarta", "Laki-laki"}
	//var animal4 *Animal = &Animal{"Burung", "jakarta", "Laki-laki"}
	var animal2 *Animal = &animal1
	var animal3 *Animal = &animal1

	animal2.Kelamin = "Perempuan"

	*animal2 = Animal{"Singa", "Kalimantan", "Laki-laki"}

	var animal4 *Animal = new(Animal)
	animal4.Kota = "Bogor"

	fmt.Println(animal1)
	fmt.Println(animal2)
	fmt.Println(animal3)
	fmt.Println(animal4)

}