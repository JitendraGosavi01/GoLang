package main
import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)


type Employee struct{
	Name string `json:"name"`
	EmpId string `json:"empId"`
	Email string `json:"email"`
}

type Employees []Employee
func main(){
   handleRequest()
}

func handleRequest(){
	myRouter:= mux.NewRouter().StrictSlash(true)
	 myRouter.HandleFunc("/", homePage)
	 myRouter.HandleFunc("/employees", getAllEmployees).Methods("GET")
	 myRouter.HandleFunc("/employees", createEmployee).Methods("POST")
	 log.Fatal(http.ListenAndServe(":8080", myRouter))
}	

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "You have hit to home page")
}

func getAllEmployees(w http.ResponseWriter, r *http.Request){
   employees:= Employees{
	   Employee{Name:"Jitendra",EmpId:"123", Email:"jg00601363@techmahindra.com"},
   }
	fmt.Println("End point hit to all employees")
	json.NewEncoder(w).Encode(employees)
}

func createEmployee(res http.ResponseWriter, req *http.Request){
   body,err := ioutil.ReadAll(req.Body)
   if err != nil {
	   panic(err.Error())
   }
 
	//fmt.Fprint(res, string(body))
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	fmt.Fprint(res, raw)
   //  out, _ := json.Marshal(body)
    //fmt.Fprint(res, string(out))
   
}
