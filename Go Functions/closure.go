package main
import "fmt"
func main(){
	num:=10
	squareNum:= func()(int){
		 num*=num
		 return num
	}
	fmt.Println("Squearing from inner function:", squareNum())
	fmt.Print("Squearing from inner function:", squareNum())
}