package main
import "fmt"
func main(){
	sum:=0
	doc:= []int{2,3,4}
	for _, num := range doc {
		sum += num
		//fmt.Println("Index is:", index, "Number is:", num);
	}
	fmt.Print("Sum is:\n", sum)
	for i, num := range doc {
		if(i == 2){
			fmt.Println("Index is:", i, "Number is:", num);
		}
		
	}

	json := map[string]string {"1":"Mango", "2":"Banana","3":"Apple"}
	for i,fruit := range json {
      fmt.Printf("%s => %s \n", i, fruit);
	}

	for i := range json {
		fmt.Println("Key => ", i)
	}
}