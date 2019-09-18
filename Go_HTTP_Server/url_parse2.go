package main
import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/employee", displayEmployee)
	http.ListenAndServe(":8000", nil)
}

func displayEmployee(res http.ResponseWriter, req *http.Request){
	//http://localhost:8000/employee?name=Jitendra&age=28&empId=123
	io.WriteString(res, "name:"+req.FormValue("name")) //name:Jitendra
	io.WriteString(res, "\nage:"+req.FormValue("age")) //age:28
	io.WriteString(res, "\nEmployeeId:"+req.FormValue("empId")) //EmployeeId:123
}