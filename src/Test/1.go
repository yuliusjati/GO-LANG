package main

import "fmt"

func main() {

	for i := 1; i <6; i++ {
		if  i&2 == 0 {
			fmt.Println("ganjil", i)
		} else  {
			fmt.Println("Genap", i)
			
		}
	}

	
}