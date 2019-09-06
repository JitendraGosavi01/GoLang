package main
import "fmt"
import "reflect"
import "strings"
/*
  strings.Split([string], [delimiter]) //returns array of string
*/
func main(){
	//string initalization
	var str string = "\t\tHELLO WORLD\t\tGO"
	fmt.Println("str:",str)
	fmt.Println("Trimmed string:",strings.TrimLeft(str, "\t"))
	fmt.Println("Trimmed string:",strings.TrimRight(str, "\t"))
	
	fmt.Println("Type:",reflect.TypeOf(str))
	
}