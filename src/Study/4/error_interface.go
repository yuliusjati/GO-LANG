package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int)(int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else {
		result := nilai/ pembagi
		return result, nil
	}
}

func main(){
	var ContohError error = errors.New("Ups error")
	fmt.Println(ContohError)

	hasil, err := pembagi(0,9)
	if err == nil{
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}