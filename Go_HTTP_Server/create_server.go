package main
import (
	"fmt"
	"net/http"
	"log"
	"html"
)
func main(){
	//http route
	http.HandleFunc("/",MyHome)
	http.HandleFunc("/about",AboutMe)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func MyHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %s\n", r)
	fmt.Fprint(w, "Hello there , thanks to visit my home")
}

func AboutMe(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	fmt.Fprint(w, "Hello there , thanks read about me")
}