/*
Struct is a data type and can be used as an anonymous field (having only the type). One struct can be inserted or "embedded" into other struct.

It is a simple 'inheritance' which can be used to implement implementations from other type or types.
*/
package main
import "fmt"
type person struct{
	fname string
	lname string
}

type employee struct{
	person
	empId int
}

func (p person) details(){
	fmt.Println(p," "+ "I am a person")
}

func (e employee) details(){
	fmt.Println(e, " "+"I am a employee")
}
func main(){
 p1:= person{"Jitendra", "Gosavi"}
 p1.details()
 e1:= employee{person:person{"Jitendra", "Gosavi"}, empId: 1234}
 e1.details()
}
