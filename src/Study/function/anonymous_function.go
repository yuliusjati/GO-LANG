package main

import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){ // ini TRUE 
		fmt.Println("you are blocked" , name)
	} else {
		fmt.Println("Welcome", name)
	}
}
//func BlacklistAdmin(name string)string{
//	return name == "admin"
//}
//func BlacklistRoot(name string)string {
//	return name == "Root"
//}

// bisa langsung dimasukan ke variable ataupun parameter tanpa memasukan nama

func main(){
	var blacklist = func(name string)bool {
		return name == "Admin"
	}

	registerUser("yulius" , blacklist)
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}