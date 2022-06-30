package onclass

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is user")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is create user")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is order")
}

func main() {
	server := NewHttpServer("test-server")
	//server.Route("/", home)
	//server.Route("/user", user)
	//server.Route("/user/create", createUser)
	//server.Route("/order", order)
	server.Route(http.MethodGet, "/user/signup", SignUp)
	server.Start("8090")
}
