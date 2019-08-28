package main
import (
	"reflect"
	"fmt"
	"strconv"
	
)

func main(){
	var a int = 10;
	var b float64 = 10.2;
    var str string = "10";
	var str1 string = "10.123"
	fmt.Println("Int to float:", float64(a), "now type of a is:", reflect.TypeOf(float64(a)))
	fmt.Println("float64 to int:", int(b), "now type of b is:", reflect.TypeOf(int(b)))

	newInt, _:= strconv.ParseInt(str, 0, 64)
	fmt.Println("string to int64:", newInt, "now type of str is:", reflect.TypeOf(newInt))

	newFloat, _:= strconv.ParseFloat(str1, 64)
	fmt.Println("string to float64:", newFloat, "now type of str1 is:", reflect.TypeOf(newFloat))
}