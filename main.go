package main

import (
	"context"
	"go_microservices/handlers" // our own package
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags) // prefix and flag
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l) // inject l in hh
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() { // handling listen and serve here not to block
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// All of the next part of our code here is used to gracefully shutdown
	// our application when the server is killed (ctrl+c, for example)
	sigChan := make(chan os.Signal) // creating a channel
	// will broadcast a message on this channel when kill or interrup happens
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// blocking here, reading from a channel will block here until there is a message available to be consumed, after consuming msg then shutdown
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 90*time.Second) //timeout context, this returns an object, with the "_" we discart it
	s.Shutdown(tc)                                                     // graceful shutdown: wait 90sec until requests are done then shutdown, good for upgrades in our app, etc

	// register handler with server
	//http.ListenAndServe(":9090", sm) // if nil is given here it uses defaultServeMux
}
