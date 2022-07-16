package main

import "fmt"
//factorial loop
func factorialloop(value int)int{
	result := 1
	for i := value; i > 0; i--{
		result *= i
	}
	return result
}

//factorial recursive
func factorialrecursive(value int)int{
	if value == 1 {
		return 1
	} else {
		return value * factorialrecursive((value-1))
	}
}
func main(){
	fmt.Println(factorialloop(5))
	fmt.Println(factorialrecursive(5 ))
}
