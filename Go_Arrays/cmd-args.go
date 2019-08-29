package main
import (
	"fmt"
	"os"
)

/*
When we need to execute a program with some arguments, 
we generally use command line argument. 
The arguments passed from the console can be received by the Go program and it can be used as an input.
*/
func main(){
	fmt.Println(len(os.Args))
    for i:=0; i< len(os.Args); i++ {
		 fmt.Println(os.Args[i])
	}
}