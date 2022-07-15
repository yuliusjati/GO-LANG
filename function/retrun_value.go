package  main

import "fmt"

func hello(nama string)string{
	return "hello " + nama
}

func bro(name string)string{
	if name == ""{
		return "hello"
	} else {
		return "halo " + name
	}
}

func main(){
	result := hello("yulius")
	fmt.Println(result)

	broo := bro("yuli")
	fmt.Println(broo)

	fmt.Println(bro(""))

}
