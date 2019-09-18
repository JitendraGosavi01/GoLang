package main
import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	//read file from the same directory
	stream, err:= ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stream) //[72 105 46 46 46 32 116 104 101 114 101 32 116 104 105 115 32 105 115 32 116 101 115 116 32 109 101 115 115 97 103 101 33]
	readString:= string(stream) //convert stream dat to string
	fmt.Println(readString) //Hi... there this is test message!
}