package controlla

import (
	"html/template"
	"hyped/viewmodel"
	"net/http"
)

type index struct {
	indexTemplate *template.Template
}

func (i index) ApplyEndpoint() {
	http.HandleFunc("/index", i.IndexHandler)
	http.HandleFunc("/home", i.IndexHandler)
	http.HandleFunc("/", i.IndexHandler)
}

func (i index) IndexHandler(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewBase()

	i.indexTemplate.Execute(w, vm)
}
