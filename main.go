package main

import (
	"context"
	"go-http-microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/gorilla/mux"
)

func main () {

	

	log := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	//Create Handlers 
	//h1 := handlers.NewHello(log)
	//g1 := handlers.NewGoodBye(log)
	ProductHandler := handlers.NewProducts(log)
	
	//Create a Server mux and register handlers using Gorilla framework 
	Gsrvmux := mux.NewRouter()
	//srvmux.Handle("/hello", h1)
	//srvmux.Handle("/Goodbye", g1)
	getRouter := Gsrvmux.Methods("Get").Subrouter()
	getRouter.HandleFunc("/",ProductHandler.GetProductshandle)
	
	putRouter := Gsrvmux.Methods("Put").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}",ProductHandler.UpdateProductshandle)

	//Create a new server 
	srv := &http.Server{
		Addr: ":8080",
		Handler: Gsrvmux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	//start the server
go func ()  {
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}()
	//trap signal or interup and gracefully showtodwn the service/server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	// Block until a signal is received.
	sig := <- sigChan
	log.Println("Received terminate singnal, graceful shutdown", sig)

	shut, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(shut)
}