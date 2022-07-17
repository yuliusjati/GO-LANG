 package main

 import "fmt"

 func main(){
	var nilai  = 67
	var rank = 3
	var rata bool = nilai > 75
	var ranking = rank <=35
	var hasil = rata || ranking

	fmt.Println(hasil)

	fmt.Println(nilai >= 60 || rank <= 30)
 }