package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	var port int

	flag.IntVar(&port, "p", 8000, "Specify the port. Default is 8000")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./html"))

	http.Handle("/", fileServer)

	fmt.Println("Webserver is now starting on port ", port)

	var portString string = ":" + strconv.Itoa(port)

	if err := http.ListenAndServe(portString, nil); err != nil {
		log.Fatal(err)
	}

}
