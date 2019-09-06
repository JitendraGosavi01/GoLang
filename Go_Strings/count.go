package main
import "fmt"
import "reflect"
import "strings"
/*
  strings.Count(string, [letter to find count])
  gives the precise count if the letter from string
*/
func main(){
	//string initalization
	var str string = "HELLO WORLD "
	var str1 string = "google.com"
	fmt.Println("str:",str)
	fmt.Println("str1:",str)
	//string finding specific occurance
	fmt.Printf("%s contains %d times L\n",str, strings.Count(str, "L")) //Is Hell present in the string: false
	fmt.Printf("%s contains %d times g\n",str1, strings.Count(str1, "g")) //Is Hell present in the string: true
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
	fmt.Println("Type:",reflect.TypeOf(str1))
}