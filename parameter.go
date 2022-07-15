package main

import "fmt"

func hello(firstname string, lastname string) {
	fmt.Println("hello", firstname, lastname)
}
func main() {
	firstname := "yulius"
	hello(firstname, "jati")
	hello("suka", "budi")
}