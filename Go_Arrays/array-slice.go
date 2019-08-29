package main
import "fmt"
func main(){
	var array = [6]int {1,2,3,4,5,6}
	fmt.Println(array);
	//slice first two elements
	fmt.Println("slice first two elements:",array[2:6])
	//slice last two elements
	fmt.Println("slice last two elements ",array[0:4])
}