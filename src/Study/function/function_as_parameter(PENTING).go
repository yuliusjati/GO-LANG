package main

import "fmt"

type nilai func(string)string

func sayhai(name string, say nilai){
	jawaban := say(name)
	fmt.Println("Selamat Sore", jawaban)
}
func jawab(name string)string{
	if name == "James"{
		return "james"
	} else {
		return " "
	}
}
func main(){
	sayhai("yulius", jawab)

	
	jwb := jawab
	sayhai("James", jwb)
}

