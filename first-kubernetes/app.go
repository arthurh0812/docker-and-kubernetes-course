package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arthurh0812/first-kubernetes/tasks"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" { // default port
		port = "8080"
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("Hi there! My name is Max"))
		if err != nil {
			log.Fatalf("error writing response: %v", err)
		}

		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	})

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			tasks.GetTasks(w, req)
		case http.MethodPost:
			c := tasks.NewTaskCreator(w, req)
			c.CreateTask()
		case http.MethodPatch:
			tasks.UpdateTask(w, req)
		}
	})

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

