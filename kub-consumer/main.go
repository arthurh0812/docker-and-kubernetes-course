package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("10.132.1.208:80/story")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	n := len(bs)
	fmt.Printf("communication worked: %d bytes transmitted\n", n)
	fmt.Println(string(bs))

	http.ListenAndServe(":80", nil)
}
