package main

import "fmt" 

func main() {
	var nama [3]string
	nama[0] = "Yulius"
	nama[1] = "Jati"
	nama[2] = "Satriyo"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])
	
	// no RIBET2
	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)

	fmt.Println(len(nama))
	fmt.Println(len(values))

	// len(array) itu panjang data bukan jumlah
}