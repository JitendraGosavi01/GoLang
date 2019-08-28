package main
import "fmt"
func main(){
	var age int
	Loop:
	fmt.Print("Please enter your age: ")
	fmt.Scan(&age)
	if(age < 18){
		fmt.Print("You can not vote, ")
		goto Loop
	}else{
		fmt.Print("You can vote")
	}

}