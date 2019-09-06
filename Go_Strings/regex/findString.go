package main
import "fmt"
import "regexp"
/**
 regexp.MustCompile([string])
*/
func main(){
	//define regular expression
	regex:= regexp.MustCompile(".com");
	/*returns specified regular expression string from provided string
	  if does not find returns nothing
	*/
	fmt.Println("String www.google.com", regex.FindString("www.google.com"))  	 //String www.google.com .com
	fmt.Println("String www.google.com", regex.FindString("www.google"))  	 //String www.google.com 
	
}