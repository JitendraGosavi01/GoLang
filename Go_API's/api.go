package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	//Employee
	Employee struct {
		Name  string `json:"name"`
		EmpId string `json:"empId"`
		Email string `json:"email"`
	}
	//Employees
	Employees []Employee
)


func main(){
   HandleRequest()
}
/*

 */
func HandleRequest(){
	 http.HandleFunc("/", HomePage)
	 http.HandleFunc("/employees", GetAllEmployees)
	 http.ListenAndServe(":8080", nil)
}


func HomePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "You have hit to home page")
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request){
   employees:= Employees{
	   Employee{Name:"Jitendra",EmpId:"123", Email:"jg00601363@techmahindra.com"},
   }
	fmt.Println("End point hit to all employees")
	json.NewEncoder(w).Encode(employees)
}
