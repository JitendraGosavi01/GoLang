package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string initalization
	var str string = "HELLO WORLD "
	fmt.Println("Default string:",str)
	//string finding specific occurance
	fmt.Println("Is H & W present in the string?",strings.ContainsAny(str, "H & W")) //Is Hell present in the string: false
	fmt.Println("Is h & w present in the string?",strings.ContainsAny(str, "h & w")) //Is Hell present in the string: true
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}