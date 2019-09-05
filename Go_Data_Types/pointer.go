//A pointer is a variable that stores the address of another variable. The general form of a pointer variable declaration is:
package main
import "fmt"
func main(){
	x:= 10;
	fmt.Println("Address of x is:",&x)
	changeX(&x)
	fmt.Println("Changed address:",x)
}
func changeX(x *int){
	*x=0
}
