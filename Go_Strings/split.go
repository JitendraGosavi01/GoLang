package main
import "fmt"
import "reflect"
import "strings"
/*
  strings.Split([string], [delimiter]) //returns array of string
*/
func main(){
	//string initalization
	var str string = "HELLO,WORLD"
	fmt.Println("str:",str)
	fmt.Println("Splitted string array:",strings.Split(str,","))
	//returns array of splitted string
	splittedString:= strings.Split(str,",");
	//loop through returned array
	for i,val:= range splittedString{
		fmt.Printf("index = %d and value = %s\n", i, val)
	}
	fmt.Println("Type:",reflect.TypeOf(str))
	
}