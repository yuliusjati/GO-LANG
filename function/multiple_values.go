package main
import "fmt"
func nameage()(string, int, string){
	return "yulius", 18, "address"
}

func main(){
	nama, umur, _ := nameage()
	fmt.Println(nama ,  umur)
	//fmt.Println(alamat)
}