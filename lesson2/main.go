package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "golang_microservices/lesson2/handler"
	// handlers "github.com/x01000101x/go-microservice/lesson2/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	//creating a server with timeouts
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//using go func because this will block the graceful shutdown
	go func() {
		//To start the server
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//Broadcast to the channel if the system is killed/interrupted
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	//Block until the messsage is received from the sigChan
	sig := <-sigChan
	l.Println("Received Terminate, graceful shutdown", sig)

	//these two lines gracefully shutdown the server
	//gracefully shutdown = shutting down the server but waiting for active connections to finish their tasks for n seconds
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
