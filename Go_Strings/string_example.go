package main
import "fmt"
import "reflect"
func main(){
	//string initalization
	var str string = "Hello world"
	fmt.Println(str)
	//find type of string
	fmt.Println(reflect.TypeOf(str))
}