package main

import "fmt"

func logging(){
	fmt.Println("selesai memanggil function")
}

func runapplication(value int){
	defer logging()
	fmt.Println("Run Applicaktion")
	result := 10 / value
	fmt.Println("result", result)
	
}

func main() {
	runapplication(0)
}