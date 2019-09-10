/**
The defer keyword is generally used for cleaning purpose. The defer keyword postpones the execution of a function or statement until the end of the calling function.

It executes code (a function or expression) when the enclosing function returns before the closing curly brace }. It is also executed if an error occurs during the execution of the enclosing function.
*/

package main
import "fmt"

func main(){
	defer Print1("hi...")
	Print2("there..")
}
func Print1(str string){
	fmt.Println(str)
}

func Print2(str string){
	fmt.Println(str)
}