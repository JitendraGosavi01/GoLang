package main
import "fmt"
type Location struct{
	lat, Long float64
}
var m = map[string]Location{
	"India":Location{40.20136, -12.3214},
	"America":Location{40.20136, -12.3214},
}
func main(){
	fmt.Println(m) //map[America:{40.20136 -12.3214} India:{40.20136 -12.3214}]
}