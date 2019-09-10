package main
//import "fmt"
import "os"

func main(){
 panic("Error Situation")
 _,Err:= os.Open("/tmp/file")
 if Err != nil{
	 panic(Err)
 }
}