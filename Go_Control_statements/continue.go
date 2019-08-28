package main

import "fmt"
func main(){
	a:= 5
	for i:=1;i<11; i++ {
		if(i == a || i == 7){
			continue
		}
		fmt.Println("Value is: ", i)
	}
}