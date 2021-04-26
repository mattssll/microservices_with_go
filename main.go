package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { // this implement the interface Handler and add to defaultServeMux
		log.Println("hello world")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}

		fmt.Fprintf(rw, "Hello, this is our data: %s", data)
	})

	http.ListenAndServe(":9090", nil) // if nil is given here it uses defaultServeMux
}
