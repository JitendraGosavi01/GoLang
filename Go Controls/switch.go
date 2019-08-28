package main
import "fmt"

func main(){
	var a int = 0;
	fmt.Print("Please enter any nuber from 1 to 5: \n")
	fmt.Scanln(&a);
	switch(a){
	case 1:
		fmt.Print("You have entered 1")
	case 2:
		fmt.Print("You have entered 2")
	case 3:
		fmt.Print("You have entered 3")
	case 4:
		fmt.Print("You have entered 4")
	case 5:
		fmt.Print("You have entered 5")
	default:
		fmt.Print("You have entered", a)
	}
}