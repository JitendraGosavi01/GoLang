package main
import "fmt"
func main(){
	var newArr [5]int
	for i:=0;i< 5; i++ {
		//push to array 
		newArr[i] = i+10 
		fmt.Printf("Emelemt [%d] = %d\n",i, newArr[i])
	}
	fmt.Print(newArr);
}