package main// recover harus didalam defer function
import "fmt"
func endApp(){
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
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