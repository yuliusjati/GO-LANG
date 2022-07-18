package main

import (
	"sort"
	"fmt"
)

func main(){
	number := [...]int{1,4,5,6,8,2}

	for i := 0; i < 6; i++ {
		bil := number[i]
		fmt.Println(bil)
		for j := 0; j < bil; j++ {
			
			fmt.Print("|")
			
		}
		fmt.Println("")
	
	}
	fmt.Println("=++++++++++++++++++++++++=")

	arrslice := number[:]
	sort.Sort(sort.IntSlice(arrslice))

	for i := 0; i < number[i]; i++ {
		bil := number[i]
		fmt.Print(bil)
		for j := 0; j < 9; j++ {
			fmt.Println("")
			fmt.Println("|")			
		}		
		fmt.Println("")
	} 
}