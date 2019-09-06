package main
import "fmt"
import "regexp"
/**
 [var name].FindStringSubmatch([string])
*/
func main(){
	//define regular expression
	regex:= regexp.MustCompile("([a-zA-Z]+)ing");
	/*returns array of string and fined regular expression string from provided string
	  if does not find returns empty array
	*/
	fmt.Println("Submatch", regex.FindStringSubmatch("Bloating"))  	 //Submatch [Bloating Bloat]
	fmt.Println("Submatch", regex.FindStringSubmatch("Floating"))  	 //Submatch [Floating Float]
	fmt.Println("Submatch", regex.FindStringSubmatch("Doing"))  	 //Submatch [Doing Do]
	fmt.Println("Submatch", regex.FindStringSubmatch("welcome"))  	 //Submatch []
	
}