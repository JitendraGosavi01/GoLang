package main
import ("fmt"
"reflect")
func main(){
	var strArr = [4]string {"Jitendra", "Gosavi", "Umesh", "Vyawhare"}
	fmt.Println(strArr); //[Jitendra Gosavi Umesh Vyawhare]
	fname1:= strArr[0:2]; 
	fname2:= strArr[2:4];
	fmt.Println(fname1); //[Jitendra Gosavi]
	fmt.Println(fname2); //[Umesh Vyawhare]
	fmt.Println(reflect.TypeOf(fname2)); // []string
	//change value fo array element
	fname1[0] = "Akshay"
	fmt.Println(fname1); //[Akshay Gosavi]

	//Omit Lower or Upper bonds
	fname1 = strArr[:2]; 
	fname2 = strArr[2:];
	fmt.Println(fname1); //[Jitendra Gosavi]
	fmt.Println(fname2); //[Umesh Vyawhare]
}