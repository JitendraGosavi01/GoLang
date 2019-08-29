package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string initalization
	var str string = "Hello world"
	fmt.Println("Default string:",str)
	//string conversion to UPPER Letters returns true or false 
	fmt.Println("Boolean:",strings.HasSuffix(str,"rld"))
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}