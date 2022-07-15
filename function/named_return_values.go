package main

import "fmt"

func fullname() (firstname, middlename, lastname string){
	firstname = "yulius"
	middlename = "jati"
	lastname = "satriyo"
	return 
}

func main(){
	firstname, middlename, lastname := fullname()
	fmt.Println(firstname, middlename, lastname)
}