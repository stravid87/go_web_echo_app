// next: crud operations https://www.youtube.com/watch?v=lf_kiH_NPvM&ab_channel=CodingwithRobby

package main

import (
	"flag"
	"fmt"
	"go-web-echo-app/config"
	"go-web-echo-app/router"
	"log"
	"net/http"
)

var (
	PORT = flag.Int("port", 3001, "Port to listen on")
)

func init() {
	config.SetupDatabaseConnection()
}

func main() {
	flag.Parse()
	fmt.Println("Starting go-web-echo-app on port: ", *PORT)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *PORT),
		Handler: http.HandlerFunc(router.Router),
	}
	log.Fatal(server.ListenAndServe())
}
