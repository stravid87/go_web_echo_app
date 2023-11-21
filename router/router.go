package router

import (
	"go-web-echo-app/controller"
	"log"
	"net/http"
)

var Router = func(w http.ResponseWriter, r *http.Request) {
	// allow from all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// routing
	switch path := r.URL.Path; {
	case path == "/":
		controller.Home(w, r)
	case path == "/page2":
		controller.PageTwo(w, r)
	case path == "/page3":
		controller.PageThree(w, r)
	case path == "/create-user":
		controller.CreateUser(w, r)
	case path == "/edit-user":
		controller.EditUser(w, r)
	case path == "/get-user":
		controller.GetUser(w, r)
	case path == "/delete-user":
		controller.DeleteUser(w, r)
	// case strings.HasPrefix(path, "/assets"):
	// 	// serve static files
	// 	http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))).ServeHTTP(w, r)
	default:
		http.Error(w, "Invalid path", http.StatusNotFound)
	}
	log.Printf("%d %s\t%s", http.StatusOK, r.Method, r.URL.Path)
}
