package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string initalization
	var str string = "Hello world"
	fmt.Println("Default string:",str)
	//string conversion to UPPER Letters
	fmt.Println("Upper case string:",strings.ToUpper(str))
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}