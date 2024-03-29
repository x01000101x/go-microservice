package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "data = %s", d)

}
