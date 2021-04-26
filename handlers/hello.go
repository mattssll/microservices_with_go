package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// create struct that implements interface http.handler
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello { // returns a hello handler (*Hello)
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world") //Println method

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return // to terminate flow of app
	}

	fmt.Fprintf(rw, "Hello, this is our data: %s", d)
}
