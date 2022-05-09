package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Starting http server...")
	done:=make(chan os.Signal,1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/healthz", healthzHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe")
	}
	fmt.Println("Server start")
	<-done
	fmt.Println("Server stop")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	req,err:=http.NewRequest("GET","http://service1",nil)
	if err != nil{
		fmt.Println(err)
	}
	nextHeader := make(http.Header)
	for key,value := range r.Header{
		nextHeader[strings.ToLower(key)] = value
	}
	req.Header = nextHeader

	client := http.Client{}
	resp,err := client.Do(req)
	if err!=nil{
		fmt.Println("request fail",err)
	}else {
		fmt.Println("request success")
	}
	if resp !=nil{
		resp.Write(w)
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
