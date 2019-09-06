package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string array initalization
	var str string = "HELLO WORLD "
	fmt.Println("Default string:",str)
	//string repeatation 3 times
	fmt.Println("Repeated string:",strings.Repeat(str, 3)) //Repeated string: HELLO WORLD HELLO WORLD HELLO WORLD
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}