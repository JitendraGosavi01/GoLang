package main
import ("fmt")
func main(){
	var numArr = []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(numArr); //[1 2 3 4 5 6 7 8 9 10]
	showInfo(numArr); // Lengh is = 10 and capacity is = 10 new array is = [1 2 3 4 5 6 7 8 9 10]
	showInfo(numArr[:2]) //Lengh is = 2 and capacity is = 10 new array is = [1 2]
	showInfo(numArr[:8]) //Lengh is = 8 and capacity is = 10 new array is = [1 2 3 4 5 6 7 8]
	showInfo(numArr[2:]) //Lengh is = 8 and capacity is = 8 new array is = [3 4 5 6 7 8 9 10]
	showInfo(numArr[:0]) //Lengh is = 0 and capacity is = 10 new array is = []
	showInfo(numArr[0:]) //Lengh is = 10 and capacity is = 10 new array is = [1 2 3 4 5 6 7 8 9 10]

	/*
	SLICE MAKE FUNCTION
	We can also create slice with the help of make function. The make function creates a zero sized array and returns slice of the array.
	*/
	slice1:= make([]int, 10 );
	sliceMake("Slice1",slice1) //Lengh is = 10 and capacity is = 10 new array Slice1 is = [0 0 0 0 0 0 0 0 0 0]
	slice2:= make([]int, 3, 10 )
	sliceMake("Slice2",slice2[:3]) //Lengh is = 3 and capacity is = 10 new array Slice2 is = [0 0 0]
	
}

func sliceMake(str string, s []int){
	fmt.Printf("Lengh is = %d and capacity is = %d new array %s is = %v\n", len(s), cap(s),str, s)
}
func showInfo(s []int){
	fmt.Printf("Lengh is = %d and capacity is = %d new array is = %v\n", len(s), cap(s), s)
}