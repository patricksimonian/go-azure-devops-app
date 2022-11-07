package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Finding mounted file!")
	fPath, ok := os.LookupEnv("CONFIG_PATH")

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	content, err := os.ReadFile(fPath)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(content)
}

func main() {
	fmt.Println("Hello world!")

	httpInvokerPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + httpInvokerPort)
	} else {
		httpInvokerPort = "3000"
	}
	http.HandleFunc("/foo", httpTriggerHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
	log.Fatal(http.ListenAndServe(":"+httpInvokerPort, nil))
}
