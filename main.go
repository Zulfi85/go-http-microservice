package main

import (
	"log"
	"net/http"
	"os"
	"go-http-microservice/handlers"
	
)

func main () {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)
	h1 := handlers.NewHello(log)
	g1 := handlers.NewGoodBye(log)
	
	srvmux := http.NewServeMux()
	srvmux.Handle("/", h1)
	srvmux.Handle("/Goodbye", g1)



	http.ListenAndServe(":8080", srvmux)
}