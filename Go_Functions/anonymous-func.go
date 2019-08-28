package main
import "fmt"

func main(){
  //anonymous function with immediate invoke function
	func (){
	fmt.Println("Hello I am printing from anonymous functions...!")
}()

//anonymous function with variable call
	print :=func (){
	fmt.Print("Hello I am printing with variable call...!")
}
print();

}

