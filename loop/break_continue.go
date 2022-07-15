package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println("perulangan ke ", i)
	}
	for i := 0; i < 10; i++{
		if i%2 == 1{
			continue
		}
		fmt.Println("perulangan ke ", i)
	}
}