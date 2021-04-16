package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var service string
var resource string

func init() {
	service = os.Getenv("SERVICE")
	if service == "" {
		service = "localhost:80"
	}
	resource = os.Getenv("RESOURCE")
	if resource == "" {
		resource = "story"
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		resp, err := http.Get(fmt.Sprintf("http://%s/%s", service, resource))
		if err != nil {
			log.Fatal(err)
		}

		bs, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		n := len(bs)
		fmt.Fprintf(w, "communication worked: %d bytes transmitted\n", n)
		fmt.Fprintln(w, string(bs))
	})

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}
