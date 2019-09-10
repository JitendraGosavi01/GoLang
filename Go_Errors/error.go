package main
import (
	"fmt"
	"errors"
	"math"
)
//Go already has pre defined error interface type 
/**
* Calculate squre root
* @param value float64
* @return float64
* @return error
*/
//   funcname(params type)(return type)
func Sqrt(value float64)(float64, error){
  if value < 0 {
	  return 0,errors.New("Math: negative number passed to Sqrt")
  }
  return math.Sqrt(value), nil
}

func main(){
	result, error := Sqrt(-49)
	if error != nil{
		fmt.Println(error) //Math: negative number passed to Sqrt
	}else{
		fmt.Println(result)
	}

	result1, error1 := Sqrt(49)
	if error1 != nil{
		fmt.Println(error1)
	}else{
		fmt.Println(result1) //7
	}
}