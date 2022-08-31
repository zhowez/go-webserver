package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

//const FSPATH = "./html/"

func main() {

	var port int
	var directory string

	flag.IntVar(&port, "p", 8000, "Specify the port. Default is 8000")
	flag.StringVar(&directory, "d", "", "Specify the sub directory to ./html where the index.html is located")
	flag.Parse()

	indexPath := "./html/" + directory + "/"

	fileServer := http.FileServer(http.Dir(indexPath))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//If the requested file exists then return if; otherwise return index.html (fileserver default page)
		if r.URL.Path != "/" {
			fullPath := indexPath + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err := os.Stat(fullPath)
			//overides the error by changing the path
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
	http.HandleFunc("/api", proxyPass)

}

func proxyPass(res http.ResponseWriter, req *http.Request) {
	// Encrypt Request here
	// ...

	url, _ := url.Parse("127.0.01:9000")
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
}
