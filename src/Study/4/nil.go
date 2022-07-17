package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
  	}
}
func main() {
	var person map[string]string = nil
	fmt.Println(person)

	data := newMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
	
	newmap := newMap("devon")
	fmt.Println(newmap)
}