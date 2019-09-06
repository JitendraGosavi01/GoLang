package main
import "fmt"
import "strings"

/*
  Compare() method: You can also compare two string using inbuilt function Compare() provide by the string package. This function return integer value after comparing two strings lexicographically. The return values are:

Return 0, if str1 == str2.
Return 1, if str1 > str2.
Return -1, if str1 < str2.

In mathematics, the lexicographic or lexicographical order (also known as lexical order, 
	dictionary order, alphabetical order or lexicographic(al) product) 
	is a generalization of the way words are alphabetically ordered based on the alphabetical order of their component letters.
*/
func main(){
	//string comparisation
	fmt.Println(strings.Compare("a","b")) //-1 if str1 < str2
	fmt.Println(strings.Compare("a","a")) //0 if str1== str2
	fmt.Println(strings.Compare("b","a")) //1 if str1 > str2 

	
}