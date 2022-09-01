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

func staticHandler(w http.ResponseWriter, r *http.Request, directory string, port int) {
	fmt.Println("home")
	indexPath := "./html/" + directory + "/"

	fileServer := http.FileServer(http.Dir(indexPath))
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
}

func apiHandler(w http.ResponseWriter, r *http.Request, port int) {
	fmt.Println("api")
	fmt.Println("reached proxy section")
	fmt.Println(r.URL)
	proxyUrl := "http://127.0.0.1:" + strconv.Itoa(port)
	fmt.Println(proxyUrl)
	url, _ := url.Parse(proxyUrl)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

const FSPATH = "./html/"

func main() {

	var port int
	var proxyPort int
	var directory string

	flag.IntVar(&port, "p", 8000, "Specify the port. Default is 8000")
	flag.IntVar(&proxyPort, "pp", 8001, "Specify the proxy port. Default is 8001")
	flag.StringVar(&directory, "d", "", "Specify the sub directory to ./html where the index.html is located")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		staticHandler(w, r, directory, port)

	})
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		apiHandler(w, r, proxyPort)

	})
	log.Fatal(http.ListenAndServe(":8000", mux))
	fmt.Println("Webserver is now starting on port ", port)
	fmt.Println("The api is running on ", proxyPort)

	var portString string = ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(portString, mux))
}

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// 	// "os"
// 	// "path"
// 	"strconv"
// 	// "strings"
// )

// func main() {

// 	http.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Println("reached proxy section")
// 		proxyUrl := "http://127.0.0.1:" + strconv.Itoa(proxyPort) + req.URL.RawPath
// 		fmt.Println(proxyUrl)
// 		url, _ := url.Parse(proxyUrl)
// 		proxy := httputil.NewSingleHostReverseProxy(url)
// 		proxy.ServeHTTP(res, req)
// 	})

// 	// indexPath := "./html/" + directory + "/"

// 	// fileServer := http.FileServer(http.Dir(indexPath))

// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	//If the requested file exists then return if; otherwise return index.html (fileserver default page)
// 	// 	if r.URL.Path != "/" {
// 	// 		fullPath := indexPath + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
// 	// 		_, err := os.Stat(fullPath)
// 	// 		//overides the error by changing the path
// 	// 		if err != nil {
// 	// 			if !os.IsNotExist(err) {
// 	// 				panic(err)
// 	// 			}
// 	// 			// Requested file does not exist so we return the default (resolves to index.html)
// 	// 			r.URL.Path = "/"
// 	// 		}

// 	// 	}

// 	// 	fileServer.ServeHTTP(w, r)
// 	// })

// 	fmt.Println("Webserver is now starting on port ", port)
// 	fmt.Println("The api is running on ", proxyPort)

// 	var portString string = ":" + strconv.Itoa(port)

// 	if err := http.ListenAndServe(portString, nil); err != nil {
// 		log.Fatal(err)
// 	}

// }
