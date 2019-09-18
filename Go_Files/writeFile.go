package main
import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main(){
	//creats new text file in same directory
	file,err:= os.Create("test.txt")
	//checks for error
     if(err != nil){
		 log.Fatal(err)
	 }
	 //writes into file
	 file.WriteString("Hi... there this is test message!")
	 file.Close() //close file
	 stream,err:= ioutil.ReadFile("test.txt")//reads file content in stream []byte formate
	 if err!=nil{
		 log.Fatal(err)
	 }
	 readString:= string(stream) //converts stream to string
	 fmt.Println(readString) //Hi... there this is test message!
}