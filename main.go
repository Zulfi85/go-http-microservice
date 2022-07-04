package main

import (
	"log"
	"net/http"
	"os"
	"go-http-microservice/handlers"
	
)

func main () {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h1 := handlers.NewHello(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", h1)

	http.ListenAndServe(":9090", sm)
}