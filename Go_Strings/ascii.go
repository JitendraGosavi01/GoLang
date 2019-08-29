package main
import "fmt"
import "strings"
func main(){
	//string initalization
	var alphabets = []string  {"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	//looping through string array
	for _,val:= range alphabets {
	fmt.Printf("ASCII value of %s is %d:\n",val, val[0])
	}
	fmt.Println("LOWER CASE")
	//looping through string array
	for _,val:= range alphabets {
	fmt.Printf("ASCII value of %s is %d:\n",strings.ToLower(val), strings.ToLower(val)[0])	
	}
}