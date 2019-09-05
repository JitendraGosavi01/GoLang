package main
import "fmt"
func main(){
	//creates empty map new map
	m:= make(map[string]int, 0)
	fmt.Println(m) //map[]
	m["j"] = 29 //setting key value pair
	m["p"] = 28 
	fmt.Println("New m:", m)
	//update key values 
	m["j"] = 30
	m["p"] = 29
	fmt.Println("New m:", m)
}