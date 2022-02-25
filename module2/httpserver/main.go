package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting http server...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe")
	}
}

func writeResponse(w http.ResponseWriter, r *http.Request, response map[string]interface{}, code int) {
	data, _ := json.Marshal(response)

	version := os.Getenv("VERSION")
	if version == "" {
		version = "null"
	}
	address := strings.Split(r.RemoteAddr, ":")

	w.Header().Set("VERSION", version)
	for key, values := range r.Header {
		for index, value := range values {
			if index == 0 {
				w.Header().Set(key, value)
			} else {
				w.Header().Add(key, value)
			}
		}
	}

	w.WriteHeader(code)
	w.Write(data)
	fmt.Println("ClientIP:", address[0], " ClientPORT:", address[1], " StatusCode:", code)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, r, map[string]interface{}{"result": "hello world"}, 200)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, r, map[string]interface{}{"result": "200"}, 200)
}
