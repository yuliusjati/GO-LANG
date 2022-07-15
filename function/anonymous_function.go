package main

import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	
}