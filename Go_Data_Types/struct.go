package main
import "fmt"
type person struct{
	fname string
	lname string
	age int
}

func main(){
  x:= person{fname:"Jitendra", lname:"Gosavi", age:29}
  fmt.Println(x)
  fmt.Println(x.fname)
}