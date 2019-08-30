package main
import "fmt"
import "regexp"
/**
 [var name].FindStringIndex([string])
*/
func main(){
	//define regular expression
	regex:= regexp.MustCompile(".com");
	/*returns array of index regular expression string from provided string
	  if does not find returns empty array
	*/
	fmt.Println("Index www.google.com", regex.FindStringIndex("www.google.com"))  	 //Index www.google.com [10 14]
	fmt.Println("Index www.abc.org", regex.FindStringIndex("www.abc.org"))  	 //Index www.google.com []
	fmt.Println("Index www.fb.com", regex.FindStringIndex("www.fb.com"))  	 //Index www.fb.com [6 10]
	
}