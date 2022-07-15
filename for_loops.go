package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke : ",counter)
		counter++
	}
	for counter1 := 20; counter1 >= 2; counter1--{
		fmt.Println("perulangan ke : ", counter1)
	}

	var slice= []string{"yulius", "jati", "satriyo", "utomo"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	var person = make(map[string]string)
	person["nama"] = "yulius"
	person["address"] = "Malang"

	for key, value := range person{
		fmt.Println("key", key, "=", value )
	}

}