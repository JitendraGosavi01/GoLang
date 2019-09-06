package main
import "fmt"
import "reflect"
import "strings"
func main(){
	//string initalization
	var str string = "HELLO WORLD "
	fmt.Println("Default string:",str)
	//string finding index of letter
	fmt.Println("The Hell is present on index number:",strings.Index(str, "L")) //The Hell is present on index number: 2
	fmt.Println("The HE is present on index number:",strings.Index(str, "HE")) //The HE is present on index number: 0
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
}