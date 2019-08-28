package main
import "fmt"
func main(){
	var num1 int = 0;
	var num2 int = 0;
	fmt.Print("Please enter any two number:\n")
	fmt.Scan(&num1)
	fmt.Scan(&num2);
	fmt.Print("Addition of tow numner's is:",num1+num2)
}