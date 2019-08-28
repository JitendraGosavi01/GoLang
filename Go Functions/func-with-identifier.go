package main
import "fmt"
type Friend struct{
	fname string
	mname string
	lname string
}

func (f Friend) fullname(){
	fmt.Print("Full name:", f.fname+" "+f.mname+" "+f.lname)
}

func main(){
   f:= Friend{"Jitendra","Rajendra","Gosavi"}
  f.fullname()
}
