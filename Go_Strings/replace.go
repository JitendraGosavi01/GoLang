package main
import "fmt"
import "reflect"
import "strings"
/*
  strings.Replace([string], [letter or world to replace by], [letter or world to replace with], [number of times])
  replace existing etter or world with provided letter or world
*/
func main(){
	//string initalization
	var str string = "HELLO WORLD "
	var str1 string = "google.com"
	fmt.Println("str:",str)
	fmt.Println("str1:",str)
	//replace specific letter or world with number of times
	fmt.Printf("New word with replacing only 2 L:%s\n",strings.Replace(str, "L", "O", 2)) //New word with replacing only 2 L:HEOOO WORLD
	fmt.Printf("New word with replacing only 3 L:%s\n",strings.Replace(str, "L", "O", 3)) //New word with replacing only 3 L:HEOOO WOROD
	fmt.Printf("New word with replacing only 1 g:%s\n",strings.Replace(str1, "g", "j",1)) //New word with replacing only 1 g:joogle.com
	fmt.Printf("New word with replacing only 2 g:%s\n",strings.Replace(str1, "g", "j",2)) //New word with replacing only 2 g:joojle.com
	fmt.Printf("New word with replacing with koogle:%s\n",strings.Replace(str1, "google", "koogle",strings.Count(str1,"google"))) //New word with replacing only 2 g:joojle.com
	//find type of string
	fmt.Println("Type:",reflect.TypeOf(str))
	fmt.Println("Type:",reflect.TypeOf(str1))
}