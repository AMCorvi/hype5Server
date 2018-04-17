package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Setup template prior to handling requests
	templates := populateTemplates()

	// Base URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		requestedFile := r.URL.Path[1:] /* remove char '/' first */

		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	})

	// Static public files
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/styles/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))

	// 	Set server to listen on port 8000
	http.ListenAndServe(":8000", nil)
}

func processTemplates() *template.Template {
	result := template.New("templates")
	const basePath = `templates`
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}


