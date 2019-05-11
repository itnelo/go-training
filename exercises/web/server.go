package web

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	ServerPort *int = flag.Int("port", 9595, "Port to listen.")
)

// Will be automatically executed.
// https://golang.org/doc/effective_go.html#init
func init() {
	flag.Parse()
}

// context
func rootHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("HTTP request received.", request)

	_, err := fmt.Fprintln(responseWriter, "This is a response from a webserver.")

	if nil != err {
		log.Println("An error has been occurred during writing a response.", err)

		var errorCode = http.StatusInternalServerError
		http.Error(responseWriter, http.StatusText(errorCode), errorCode)
	}
}

func HttpServerStart() {
	log.Println("Starting a webserver instance...")
	log.Println("Port to listen:", *ServerPort)

	// Adding new pattern for default router http.ServeMux.
	http.HandleFunc("/", rootHandler)

	// Serving request via default router http.ServeMux.
	var addr = ":" + strconv.Itoa(*ServerPort)
	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
