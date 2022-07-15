package main

import "fmt"

func main(){
	months := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(cap(months[6:9]))

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Mei lagi"
	//fmt.Println(months)
	
	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "yulius")
	fmt.Println(slice3)
	slice3[1] = "desember lewat"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//jadi "appent" akan membuat array baru jika array sudah tidak dapat menampung lagi
	//tapi jika array masih bisa menampung maka akan tetap menggunakan array yang lama

	newslice := make([]string, 2, 5)

	newslice[0] = "Yulius"
	newslice[1] = "Jati"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	copyslice := make([]string, len(newslice), cap(newslice))
	copy(copyslice, newslice)

	fmt.Println(copyslice)

	iniarray := [...]int{1,2,3,4,5}// bisa ... atau angka
	inislice := []int{1,2,3,4,5}

	fmt.Println(iniarray)
	fmt.Println(inislice)
}
