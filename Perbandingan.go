package main

import "fmt"

func main() {
	var nama1 = "yulius"
	var nama2 = "jati"

	var result bool = nama1 != nama2
	fmt.Println(result)

	var value1 = 70
	var value2 = 120

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <=  value2)
}