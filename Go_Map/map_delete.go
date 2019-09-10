package main
import "fmt"
func main(){
	//creates empty map new map
	m:= make(map[string]int, 0)
	fmt.Println(m) //map[]
	m["j"] = 29 //setting key value pair
	m["p"] = 28 
	fmt.Println("New m:", m) //New m: map[j:29 p:28]
	//update key values 
	m["j"] = 30
	m["p"] = 29
	fmt.Println("New mapped m:", m) //New mapped m: map[j:30 p:29]
	//delete key value from map
	//syntax = delete(mapname, keyname)
	delete(m, "j")
	fmt.Println("New mapped m after deletetion:", m) //New mapped m after deletetion: map[p:29]
}