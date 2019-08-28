package main
import "fmt"
func main(){
   fmt.Println(factorial(5))
}

func factorial(num int)int{
	if(num == 0){
		return 1
	}
	fact:= num*factorial(num-1)
	return fact;
}