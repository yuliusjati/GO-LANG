package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "mr. " + man.Name
	fmt.Println("di Method", man.Name)
}

func main() {
	yulius := Man{"yulius"}
	yulius.Married()

	fmt.Println(yulius.Name)
}