package main
import "fmt"
import "reflect"
func main(){
	//string initalization
	var str string = "Hello world"
	fmt.Println("String:",str)
	//find length of string
	fmt.Println("Length:",len(str))
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}