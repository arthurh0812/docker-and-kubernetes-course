package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var defaultGoal = "Learn Docker!"

var goal = defaultGoal

var html = `
<html>
    <head>
        <link rel="stylesheet" href="/public/styles.css">
    </head>
    <body>
        <section>
            <h2>My Course Goal:</h2>
            <h3>{{.Goal}}</h3>
        </section>
        <form action="/store-goal" method="POST">
            <div class="form-control">
                <label>Course Goal</label>
                <input type="text" name="goal" value="{{.DefaultGoal}}">
            </div>
            <button>Set Course Goal</button> 
        </form>
    </body>
</html>
`

func main() {
	// static file server
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		templ, err := template.New("base").Parse(html)
		if err != nil {
			log.Fatal(err)
		}
		data := struct {
			Goal        string
			DefaultGoal string
		}{
			Goal:        goal,
			DefaultGoal: defaultGoal,
		}
		templ.Execute(w, data)
	})

	http.HandleFunc("/store-goal", func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			http.Error(w, "failed to parse form data", http.StatusInternalServerError)
			return
		}
		inputGoal := req.PostFormValue("goal")
		fmt.Println(inputGoal)
		// overwrite the global variable "goal"
		goal = inputGoal
		http.Redirect(w, req, "/", http.StatusSeeOther)
	})

	http.ListenAndServe(":80", nil)
}
