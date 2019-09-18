/*
Go has good support for url parsing. URL contains a scheme, authentication info, host, port, path, query params, and query fragment. we can parse URL and deduce what are the parameter is coming to the server and then process the request accordingly.

The net/url package has the required functions like Scheme, User, Host, Path, RawQuery etc
*/

package main
import  ("fmt"
"net"
"net/url"
)
func main(){
	p:=fmt.Println
	s:="Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"
	//s:= "https://golang.org/pkg/net/http/#example_ListenAndServe"
	u, err:= url.Parse(s) 
	if err != nil{
		panic(err)
	}
	p(u.Scheme) //prints schema of url = mysql
	p(u.User) //prints the parsed user and URL = admin:password
	p(u.User.Username()) //prints user's name = admin
	pass,_:=u.User.Password() 
	p(pass) //prints user password = password
	p(u.Host) //prints host and port = serverhost.com:8081
	host,port,_:=net.SplitHostPort(u.Host)
	p(host) //prints host = serverhost.com 
	p(port) //prints port = 8081
	p(u.Path) //prints server path / endpoint = /server/page1
	p(u.Fragment) //prints fragment path value = X
	p(u.RawQuery) //prints the querystring/queryparamater ket value = key=value&key2=value2
	m,_:=url.ParseQuery(u.RawQuery) //parse query params into map 
	p(m) //prints param map = map[key:[value] key2:[value2]]
	p(m["key2"][0]) //prints specific key value =value2
}