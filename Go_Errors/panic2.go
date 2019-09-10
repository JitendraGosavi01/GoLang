package main
import "fmt"

func main(){
	fmt.Println("Calling x from main..")
	x()
	fmt.Println("Returned from x")

}

func x(){
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Recovered in x", r)
		}
	}()
	fmt.Println("Execiting x...")
	fmt.Println("Calling y...")
	y(0)
	fmt.Println("Returned normally from y.")
}

func y(i int){
  fmt.Println("Executing y...")
  if i>2{
	  fmt.Println("panicking...")
	  panic(fmt.Sprintf("%v", i))
  }
  defer fmt.Println("Defer in y...")
  fmt.Println("Printing in y...", i)
  y(i+1)
}

/*
Calling x from main..
Execiting x...
Calling y...
Executing y...
Printing in y... 0
Executing y...
Printing in y... 1
Executing y...
Printing in y... 2
Executing y...
panicking...
Defer in y...
Defer in y...
Defer in y...
Recovered in x 3
Returned from x
*/