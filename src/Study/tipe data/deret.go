package main

import "fmt"

func main(){
	batas := 10

	var deret = 1
	for i := 0; i <= batas; i++{
		fmt.Println(deret, "")
		deret = deret % 2
	}
}