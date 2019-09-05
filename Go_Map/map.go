package main
import "fmt"
//syntax = var var_name map[type] string/collection/json
func main(){
	x:= map[string]int {"J":29, "P":25}
	fmt.Println(x); //map[J:29 P:25]
	fmt.Println("Value of P:",x["P"]);//Value of P: 25
}