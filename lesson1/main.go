package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//Create a Http Request Handler func (curl -v localhost:9090)
	//http.ResponseWriter = Interface used by http handler to construct Http Response
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	//Create a Http Request Handler func (curl -v localhost:9090/goodbye)
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	//Create a Http Request Handler func (curl -v -d "Value" localhost:9090/test)
	http.HandleFunc("/test", func(resp http.ResponseWriter, req *http.Request) {
		log.Println("Hello World")

		//Read the data from the request to the body
		d, _ := ioutil.ReadAll(req.Body)

		//Request to the server
		log.Printf("Data = %s \n", d)
	})

	//Create a Http Request Handler func (curl -d "Value" localhost:9090/test2)
	http.HandleFunc("/test2", func(resp http.ResponseWriter, req *http.Request) {

		//Read the data from the request to the body
		d, err := ioutil.ReadAll(req.Body)

		//error handler
		if err != nil {
			http.Error(resp, "Oops", http.StatusBadRequest)
			return
		}

		//Response to the client
		fmt.Fprintf(resp, "Data = %s \n", d)
	})

	//Create Server on port 9090 and default request handler (nil)
	http.ListenAndServe(":9090", nil)
}
