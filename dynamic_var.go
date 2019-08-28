package main
import "fmt"

func main(){
	x := 10;
	y := 10.2
	z := true
	fmt.Printf("Type of x is %T and value is %d\n", x,x)
	fmt.Printf("Type of y is %T and value is %f\n", y,y)
	fmt.Printf("Type of z is %T and value is %t", z,z)
}