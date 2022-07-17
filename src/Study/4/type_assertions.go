package main

import"fmt"

func random() interface{} {
	return "KO"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)
// jika ingin aman lebih baik menggunakan switch

	switch value := result.(type){
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("unknow type") 
	}
}