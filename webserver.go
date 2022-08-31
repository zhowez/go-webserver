package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

//const FSPATH = "./html/"

func main() {

	var port int

	flag.IntVar(&port, "p", 8000, "Specify the port. Default is 8000")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./html"))

	//http.Handle("/", fileServer)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//If the requested file exists then return if; otherwise return index.html (fileserver default page)
		if r.URL.Path != "/" {
			fullPath := "./html/" + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err := os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					panic(err)
				}
				// Requested file does not exist so we return the default (resolves to index.html)
				r.URL.Path = "/"
			}

		}

		fileServer.ServeHTTP(w, r)
	})

	fmt.Println("Webserver is now starting on port ", port)

	var portString string = ":" + strconv.Itoa(port)

	if err := http.ListenAndServe(portString, nil); err != nil {
		log.Fatal(err)
	}

}
