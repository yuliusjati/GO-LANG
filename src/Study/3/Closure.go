package main

import "fmt"

func main(){
	nama := "yulius"
	number := 0

	increment := func(){
		nama := "jati"// jika tidak mau merubah "nama" di luar func harus membuat variable baru

		fmt.Println("Halo")
		number++
		fmt.Println(nama)
	}

	increment()
	increment()
	increment()

	fmt.Println(number)
	fmt.Println(nama)
}