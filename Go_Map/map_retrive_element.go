package main
import "fmt"
func main(){
	//creates empty map new map
	m:= make(map[string]int, 0)
	fmt.Println(m) //map[]
	m["j"] = 29 //setting key value pair
	m["p"] = 28 
	fmt.Println("New m:", m) //New m: map[j:29 p:28]
	//If key is not present, then the value of elem is the default value of element type.
	v, iskeyPresent := m["j"]
	fmt.Println("The value is:", v, "Present?", iskeyPresent);
	v1, iskeyPresent1:= m["k"]
	fmt.Println("The value is:", v1, "Present?", iskeyPresent1);
}