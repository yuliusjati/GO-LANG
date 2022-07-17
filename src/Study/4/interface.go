package main

import "fmt"

type hasName interface {
	getName() string
}

func sayHello(hasname hasName) {
	fmt.Println("hello", hasname.getName())
}
//implementasi
type Person struct{
	Name string
}
func (person Person)getName() string{
	return person.Name
}

type Animal struct{
	Name string
}
func (animal Animal)getName()string{
	return animal.Name
}
func main(){
	var halo Person
	halo.Name = "fadil"
	sayHello(halo)

	hewan := Animal{
		Name: "anjing",
	}
	sayHello(hewan)
}