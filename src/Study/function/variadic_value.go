package main

import "fmt"

func number(number ...int)int{
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}
func main(){
	total := number(36, 45, 67)
	fmt.Println(total)

	slice := []int{10, 20, 30}
	total = number(slice...)
	fmt.Println(total)
}