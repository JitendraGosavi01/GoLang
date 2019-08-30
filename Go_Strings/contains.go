package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string initalization
	var str string = "HELLO WORLD "
	fmt.Println("Default string:",str)
	//string finding specific occurance
	fmt.Println("Is Hell present in the string?",strings.Contains(str, "Hell")) //Is Hell present in the string: false
	fmt.Println("Is Hell present in the string?",strings.Contains(str, strings.ToUpper("Hell"))) //Is Hell present in the string: true
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}