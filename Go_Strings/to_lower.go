package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string initalization
	var str string = "HELLO WORLD"
	fmt.Println("Default string:",str)
	//string conversion to UPPER Letters
	fmt.Println("Lower case string:",strings.ToLower(str))
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}