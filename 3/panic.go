package main
import "fmt"
func endApp(){
	fmt.Println("Aplikasi selesai")
}
func runApp(errar bool){
	defer endApp()
	if errar{
		panic("aplikasi ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
func main() {
	runApp(true)
}