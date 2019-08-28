package main
import (
	"fmt" 
//"reflect"
)
func main(){
	x,y:= addAll(1,2,3,4,5,6,7,8,9,10)
	fmt.Print(x, y)
}
func addAll(args...int)(int,int){
	finalAdd:=0
	finalSub:=0;
	for _,value:= range args{
		finalAdd += value
		finalSub -= value
	}
	return finalAdd,finalSub
}