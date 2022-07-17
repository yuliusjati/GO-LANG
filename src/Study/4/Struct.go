package main

import "fmt"

type customer struct {
	name, address string
	age           int
	//school bool
}
func sayHi(Customer customer, Name string){//struct func/method
// bisa begini juga
// func (Customer customer) sayHi( Name string){
	fmt.Println("hello", Name, "my Name is", Customer.name)
}


func main() {

	var yulius customer
	yulius.name = "yulius jati satriyo utomo"
	yulius.address = "Malang"
	yulius.age = 18

	//struct func/method
	/**func (Customer customer) sayHi( Name string){
		jika kalian menggunakan yang ini maka cara ngerunnya juga beda
		kalian bisa mengunakan 
		yulius.sayHi("devon") 
	*/
	//sesuai dengan tata letak
	sayHi(yulius,"devon")

	fmt.Println(yulius)
	fmt.Println(yulius.name)
	fmt.Println(yulius.address)
	fmt.Println(yulius.age)

	fadil := customer{
		name: "fadlillah",
		address: "indonesia",
		age: 18,
	}
	fmt.Println(fadil)

	kevin := customer{"kevin", "Malang", 18}
	fmt.Println(kevin)
}