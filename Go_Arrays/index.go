package main
import "fmt"
func main(){
	var newArr = [5]int{1,2,3,4,5}
	for i:=0;i< len(newArr); i++ {
		fmt.Printf("Emelemt [%d] = %d\n",i, newArr[i])
	}
}