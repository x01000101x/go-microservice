package handlers

import (
	"log"
	"net/http"
)

//Struct Logger
type Goodbye struct {
	l *log.Logger
}

//Func to return to Goodbye Struct
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

//Func Serve HTTP
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Byee"))
}
